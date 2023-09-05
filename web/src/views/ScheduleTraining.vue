<script setup lang="ts">
import AppLayout from '@/layouts/App.vue'
import {computed, ref} from 'vue'
import router from '@/router'
import {getTrainingBalance} from '@/repositories/user'
import type {Hour, ModelDate} from '@/repositories/clients/trainer'
import {getAvailableDates, rescheduleTraining, scheduleTraining} from '@/repositories/trainings'
import {ElNotification} from 'element-plus'
import 'element-plus/es/components/notification/style/css'
import dayjs from 'dayjs'
import {TRAINING_TIMEZONE} from '@/date'

const props = defineProps<{
  isPropose?: boolean
}>()

const trainingToReschedule = computed<string | undefined>(() => {
  const param = router.currentRoute.value.params['trainingID']
  if (!param) return undefined
  if (typeof param !== 'string') return undefined
  return param.toString()
})
const trainingBalance = ref<number>(0)
const trainingData = ref<{
  date: string,
  hour: string,
  notes: string
}>({
  date: '',
  hour: '',
  notes: '',
})
const calendar = ref<Array<ModelDate>>([])
const showLoader = ref<boolean>(false)

getAvailableDates().then((data) => calendar.value = data)

getTrainingBalance().then((balance) => trainingBalance.value = balance)
    .catch((error) => console.error(error))

const availableHours = computed<Array<Hour>>(() => {
  const currentDate = calendar.value.find(obj => obj.date === trainingData.value.date)

  if (!currentDate) {
    return []
  }

  return currentDate.hours.filter(obj => obj.available === true)
})

const scheduleNewTraining = (): void => {
  showLoader.value = true

  if (trainingToReschedule.value) {
    rescheduleTraining(
        trainingToReschedule.value,
        trainingData.value.notes,
        trainingData.value.date,
        trainingData.value.hour,
        props.isPropose,
    )
        .then(() => {
          if (props.isPropose) {
            ElNotification.success({
              message: 'Training reschedule proposal sent!'
            })
          } else {
            ElNotification.success({
              message: 'Training rescheduled!'
            })
          }
          showLoader.value = false
          router.push({name: 'trainingsList'})
        })
        .catch((err) => {
          showLoader.value = false
          ElNotification.error({
            message: 'Failed to reschedule training'
          })
          console.error(err)
        })
  } else {
    scheduleTraining(trainingData.value.notes, trainingData.value.date, trainingData.value.hour)
        .then(() => {
          showLoader.value = false
          ElNotification.success({
            message: 'Training added!',
          })
          router.push({name: 'trainingsList'})
        })
        .catch((error) => {
          showLoader.value = false
          ElNotification.error({
            message: 'Failed to add training',
          })
          console.error(error)
        })
  }
}
</script>

<template>
  <app-layout>
    <div class="py-5 text-center">
      <p style="font-size: 59px;">&#128170;</p>
      <h2>
        <span v-if="!trainingToReschedule">Schedule training</span>
        <span v-if="trainingToReschedule && !isPropose">Re-schedule training {{ trainingToReschedule }}</span>
        <span v-if="trainingToReschedule && isPropose">Propose re-schedule training {{ trainingToReschedule }}</span>
      </h2>

      <br>
      <p class="lead">Below is an example form built entirely with Bootstrap’s form controls. Each required form
        group
        has a validation state that can be triggered by attempting to submit the form without completing it.</p>

      <div class="alert alert-warning" role="alert" v-if="isPropose">
        It's less than 24h left until the training. Proposition of re-schedule may be refused.
      </div>

      <div v-if="!isPropose">
        <span class="trainings-left">Trainings left: <b>{{ trainingBalance }}</b></span>
      </div>
    </div>
    <div class="row justify-content-md-center">
      <div class="col-md-8 order-md-1 l-md">
        <form class="needs-validation" @submit.prevent="scheduleNewTraining" novalidate>
          <div class="row">
            <div class="col-md-6 mb-3">
              <label for="day">Day</label>
              <select class="form-select" size="7" id="day" v-model="trainingData.date">
                <option v-for="day in calendar" :key="day.date"
                        :value="day.date">
                  {{ day.date }}
                </option>
              </select>
            </div>
            <div class="col-md-6 mb-3">
              <label for="hour">Hour</label>
              <select class="form-select" size="7" id="hour" v-model="trainingData.hour">
                <template v-for="hour in availableHours">
                  <option :key="dayjs(hour.hour).utc().format('HH:mm')"
                          :value="dayjs(hour.hour).utc().format('HH:mm')"
                          v-if="!hour.hasTrainingScheduled">
                    {{ dayjs(hour.hour).tz(TRAINING_TIMEZONE).format('HH:mm') }}
                  </option>
                </template>
              </select>
            </div>
          </div>

          <div class="row">
            <div class="col-12">
              <label for="notes">Notes <small>(visible for trainer)</small></label>
              <textarea class="form-control" id="notes" rows="3" v-model="trainingData.notes"
                        maxlength="1000"></textarea>
            </div>
          </div>

          <hr class="mb-4">
          <div class="d-grid">
            <button class="btn btn-primary btn-lg btn-block" type="submit">
              Schedule training
              <span v-if="showLoader" class="spinner-grow spinner-grow-sm"
                    style="width: 1.3rem; height: 1.3rem;"
                    role="status" aria-hidden="true"></span>
            </button>
          </div>
        </form>
      </div>
    </div>
  </app-layout>
</template>

<style scoped>
.trainings-left {
  font-size: 1.5rem;
}
</style>