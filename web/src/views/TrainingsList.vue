<script setup lang="ts">
import AppLayout from '@/layouts/App.vue'
import {computed, h, onMounted, ref, type VNode} from 'vue'
import {getUserRole, Trainer} from '@/repositories/user'
import {approveReschedule, cancelTraining, getCalendar, rejectReschedule} from '@/repositories/trainings'
import {type Training} from '@/repositories/clients/trainings'
import dayjs from 'dayjs'
import utc from 'dayjs/plugin/utc'
import timezone from 'dayjs/plugin/timezone'
import {TRAINING_TIMEZONE} from '@/date'
import {ElMessageBox, ElNotification} from 'element-plus'
import 'element-plus/es/components/message-box/style/css'
import 'element-plus/es/components/notification/style/css'

dayjs.extend(utc)
dayjs.extend(timezone)

const userRole = ref<string>('')
const isTrainer = computed<boolean>(() => userRole.value === Trainer)
const calendar = ref<Array<Training>>([])

onMounted(() => {
  userRole.value = getUserRole()
  getCalendar()
      .then((response) => calendar.value = response.data.trainings)
      .catch((error) => console.error(error))
})

const handleCancelTraining = (event: Event): void => {
  if (!(event.target instanceof HTMLElement)) return

  const target = event.target as HTMLElement
  const trainingUUID = target.getAttribute('data-training-uuid')
  if (!trainingUUID) return

  const training = calendar.value.find((t) => t.uuid === trainingUUID)
  if (!training) return

  const title = 'Are you sure you want to cancel training?'

  let message: string | VNode
  if (!training.canBeCancelled) {
    message = h('b', 'It\'s less than 24h before training, so you will not receive your credits back.')
  } else {
    message = 'Your training balance will be returned.'
  }

  ElMessageBox.confirm(message, title)
      .then(() => {
        return cancelTraining(trainingUUID)
      }, () => {
        ElNotification.error({
          message: 'Failed to cancel training'
        })
      })
      .then(() => {
        ElNotification.info({
          message: 'Training cancelled'
        })
      })
      .then(() => {
        return getCalendar()
      })
      .then((data) => calendar.value = data.data.trainings)
      .catch((err) => console.error('cancel fail', err))
}

const handleAcceptReschedule = (event: Event): void => {
  if (!(event.target instanceof HTMLElement)) return

  let target: HTMLElement | null = event.target as HTMLElement
  let trainingUUID: string | null = null
  while (target) {
    trainingUUID = target.getAttribute('data-training-uuid')
    if (trainingUUID) break
    target = target.parentElement
  }
  if (!trainingUUID) return
  const uuid = trainingUUID

  ElMessageBox.confirm(undefined, 'Are you sure you want to accept?').then(() => {
    approveReschedule(uuid).then(async () => {
      calendar.value = (await getCalendar()).data.trainings
      ElNotification.info({
        message: 'Reschedule accepted'
      })
    }).catch((err) => {
      ElNotification.error({
        message: 'Failed to accept reschedule'
      })
      console.error(err)
    })
    ElNotification.info()
  }).catch(() => {
    console.log('Clicked on cancel')
  })
}

const handleRejectReschedule = (event: Event): void => {
  if (!(event.target instanceof HTMLElement)) return

  let target: HTMLElement | null = event.target as HTMLElement
  let trainingUUID: string | null = null
  while (target) {
    trainingUUID = target.getAttribute('data-training-uuid')
    if (trainingUUID) break
    target = target.parentElement
  }
  if (!trainingUUID) return
  const uuid = trainingUUID

  ElMessageBox.confirm(undefined, 'Are you sure you want to reject')
      .then(() => {
        rejectReschedule(uuid)
            .then(() => {
              getCalendar().then((data) => calendar.value = data.data.trainings)
              ElNotification.info({
                message: 'Reschedule rejected'
              })
            })
            .catch((err) => {
              ElNotification.error({
                message: 'Failed to reject reschedule'
              })
              console.error(err)
            })
      })
      .catch(() => {
        console.log('Clicked on cancel')
      })
}
</script>

<template>
  <app-layout>
    <div class="py-5 text-center">
      <h2>Your trainings</h2>
      <p class="lead">Below is an example form built entirely with Bootstrap’s form controls. Each required form
        group
        has a validation state that can be triggered by attempting to submit the form without completing it.</p>
    </div>
    <br><br>
    <table class="table table-striped table-hover">
      <thead>
      <tr>
        <th scope="col">#</th>
        <th scope="col">When</th>
        <th scope="col">Notes</th>
        <th scope="col" v-if="isTrainer">Attendee</th>
        <th scope="col">Actions</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="(training, idx) in calendar" :key="training.uuid"
          :class="{'table-info': training.requireRescheduleApproval}">
        <th scope="row">{{ idx + 1 }}</th>
        <td>
          <span :class="{'old-date': training.proposedTime}">{{
              dayjs(training.time).tz(TRAINING_TIMEZONE).format('YYYY-MM-DD HH:mm')
            }}</span>
          <span v-if="training.proposedTime"
                :title="'proposed by ' + training.moveProposedBy"><br>{{
              dayjs(training.proposedTime).tz(TRAINING_TIMEZONE).format('YYYY-MM-DD HH:mm')
            }}</span>
        </td>
        <th>{{ training.notes }}</th>
        <th v-if="isTrainer">{{ training.user }}</th>
        <td>
          <button type="button"
                  :class="training.canBeCancelled ? 'btn btn-warning' : 'btn btn-danger'"
                  :title="training.canBeCancelled ? 'Your training balance will be returned' : 'Your training balance will be not returned because it\s less than 24h before training'"
                  @click="handleCancelTraining"
                  :data-training-uuid="training.uuid"
          >
            Cancel
          </button>
          &nbsp;
          <router-link
              :to="{ name: 'proposeNewDate', params: { trainingID: training.uuid }}"
              v-if="training.moveRequiresAccept"
          >
            <button class="btn btn-info">Propose new time</button>
          </router-link>
          &nbsp;
          <router-link v-if="!training.moveRequiresAccept"
                       :to="{ name: 'rescheduleTraining', params: { trainingID: training.uuid }}"
          >
            <button class="btn btn-primary">Move</button>
          </router-link>

          <div v-if="training.proposedTime">
            <br>
            <button type="button" class="btn btn-warning" @click="handleAcceptReschedule"
                    :data-training-uuid="training.uuid"
                    v-if="userRole !== training.moveProposedBy"
            >
              Approve reschedule
            </button>
            &nbsp;
            <button type="button" class="btn btn-warning" @click="handleRejectReschedule"
                    :data-training-uuid="training.uuid"
            >
              <span v-if="userRole !== training.moveProposedBy">Reject reschedule</span>
              <span v-if="userRole === training.moveProposedBy">Cancel reschedule request</span>
            </button>
          </div>
        </td>
      </tr>
      </tbody>
    </table>
  </app-layout>
</template>

<style scoped>
h3 {
  margin: 40px 0 0;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: #42b983;
}

body {
  background-color: #f5f5f5;
}

.old-date {
  text-decoration: line-through;
  opacity: 0.5;
}
</style>
