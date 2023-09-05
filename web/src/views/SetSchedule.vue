<script setup lang="ts">
import AppLayout from '@/layouts/App.vue'
import {getPeriods, getSchedule, type Period, setHourAvailability} from '@/repositories/trainings'
import {ref} from 'vue'
import type {Hour, ModelDate} from '@/repositories/clients/trainer'
import {ElNotification} from 'element-plus'
import 'element-plus/es/components/notification/style/css'
import dayjs from 'dayjs'
import utc from 'dayjs/plugin/utc'
import timezone from 'dayjs/plugin/timezone'
import {TRAINING_TIMEZONE} from '@/date'

dayjs.extend(utc)
dayjs.extend(timezone)

const periods = ref<Period[]>([])
const selectedPeriod = ref<string>('')
const selectedDateFrom = ref<string>('')
const selectedDateTo = ref<string>('')
const schedule = ref<ModelDate[]>([])

const selectedPeriodValue = (period: Period): string => period.from + '/' + period.to

periods.value = getPeriods()
selectedPeriod.value = selectedPeriodValue(periods.value[0])
selectedDateFrom.value = periods.value[0].from
selectedDateTo.value = periods.value[0].to
getSchedule(selectedDateFrom.value, selectedDateTo.value)
    .then((dates) => schedule.value = dates)

const changedPeriod = (event: Event): void => {
  if (!(event.target instanceof HTMLElement)) return

  const target = event.target as HTMLElement
  const from = target.getAttribute('data-from')
  if (!from) return
  const to = target.getAttribute('data-to')
  if (!to) return

  selectedDateFrom.value = from
  selectedDateTo.value = to

  getSchedule(from, to).then((dates) => schedule.value = dates)
}

const selectAllInDay = (event: Event): void => {
  if (!(event.target instanceof HTMLElement)) return

  const target = event.target as HTMLElement
  const date = target.getAttribute('data-date')
  if (!date) return

  const scheduleClone = schedule.value
  for (let scheduleIdx in scheduleClone) {
    const day = scheduleClone[scheduleIdx]

    if (date != day.date) {
      continue
    }

    const updates: string[][] = []

    for (let idx in day.hours) {
      let d = day.hours[idx].hour
      updates.push([dayjs(d).utc().format('YYYY-MM-DD'), dayjs(d).utc().format('HH:mm')])
    }

    setHourAvailability(updates, true).then(() => {
      return getSchedule(selectedDateFrom.value, selectedDateTo.value)
    }).then((dates) => schedule.value = dates)
        .catch((err) => console.error(err))
  }

  // schedule.value = scheduleClone
}

const toggleHour = (event: Event, hour: Hour): void => {
  if (!(event.target instanceof HTMLInputElement)) return

  const target = event.target as HTMLInputElement

  const hourDayjs = dayjs(hour.hour).utc()
  const updates = [[hourDayjs.format('YYYY-MM-DD'), hourDayjs.format('HH:mm')]]
  setHourAvailability(updates, target.checked)
      .then(() => {
      })
      .catch((err) => {
        console.error(err)
        hour.available = false
        ElNotification.error({
          message: 'Failed to update schedule',
        })
      })
}
</script>

<template>
  <app-layout>
    <div class="py-5 text-center">
      <p style="font-size: 59px;">&#128170;</p>
      <h2>
        <span>Set schedule</span>
      </h2>

      <br>
      <p class="lead">Below is an example form built entirely with Bootstrap’s form controls. Each required form
        group
        has a validation state that can be triggered by attempting to submit the form without completing it.</p>

      <div class="btn-group btn-group-toggle">
        <label class="btn btn-outline-dark" v-for="period in periods" :key="period.from+period.to"
               :class="{ 'active': selectedPeriod === selectedPeriodValue(period) }">
          <input type="radio" name="options" :data-from="period.from" :data-to="period.to"
                 @change="changedPeriod" v-model="selectedPeriod" :value="selectedPeriodValue(period)">
          {{ period.from }} - {{ period.to }}
        </label>
      </div>
    </div>

    <form class="needs-validation" novalidate>
      <div class="row">
        <div class="text-center schedule-column" v-for="day in schedule" :key="day.date">
          <h4 class="mb-3">{{ day.date }}</h4>
          <button type="button" class="btn btn-outline-primary btn-sm" :data-date="day.date"
                  @click="selectAllInDay">Select all
          </button>
          <br><br>

          <div v-for="hour in day.hours" :key="dayjs(hour.hour).utc().format('HH:mm')">
            <div class="btn-group-toggle" data-toggle="buttons"
                 :title="hour.hasTrainingScheduled ? 'Training scheduled on this date' : ''">
              <label
                  :class="{'btn btn-lg': true, 'active': hour.available, 'btn-primary': !hour.hasTrainingScheduled, 'btn-secondary': hour.hasTrainingScheduled}">
                <input type="checkbox" autocomplete="off" v-model="hour.available"
                       @change.prevent="toggleHour($event, hour)"
                       :disabled="hour.hasTrainingScheduled">
                {{ dayjs(hour.hour).tz(TRAINING_TIMEZONE).format('HH:mm') }}
              </label>
            </div>
            <br>
          </div>
        </div>
      </div>
    </form>

  </app-layout>
</template>

<style scoped>
.btn-primary:not(.checked) {
  color: #007bff;
  background-color: transparent;
  background-image: none;
  border-color: #007bff;
}

.schedule-column {
  width: 140px;
}
</style>