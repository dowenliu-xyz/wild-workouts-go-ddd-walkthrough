<script setup lang="ts">
import {onMounted, ref} from 'vue'
import AppLayout from '@/layouts/App.vue'
import FullCalendar from '@fullcalendar/vue3'
import {type CalendarOptions, type EventInput} from '@fullcalendar/core'
import interactionPlugin from '@fullcalendar/interaction'
import dayGridPlugin from '@fullcalendar/daygrid'
import timeGridPlugin from '@fullcalendar/timegrid'
import listPlugin from '@fullcalendar/list'
import {TRAINING_TIMEZONE} from '@/date'
import dayjs from 'dayjs'
import utc from 'dayjs/plugin/utc'
import timezone from 'dayjs/plugin/timezone'
import {getCalendar, getSchedule} from '@/repositories/trainings'
import {getUserRole, Trainer} from '@/repositories/user'

dayjs.extend(utc)
dayjs.extend(timezone)

const options = ref<CalendarOptions>({
  plugins: [interactionPlugin, dayGridPlugin, timeGridPlugin, listPlugin],
  initialView: 'timeGridWeek',
  headerToolbar: {
    left: 'prev,next today',
    center: 'title',
    right: 'dayGridMonth,timeGridWeek,timeGridDay,listWeek'
  },
  navLinks: true,
  events: [],
  timeZone: TRAINING_TIMEZONE,
})

const getScheduleCalenderEvents = async (): Promise<EventInput[]> => {
  const start = dayjs().utc().set('hour', 0).set('minute', 0).set('second', 0).set('millisecond', 0)
  const end = start.add(3, 'month')

  return getSchedule(start.format(), end.format()).then((schedule) => {
    const scheduleEvents = []
    for (let idx in schedule) {
      const date = schedule[idx]

      for (let idx in date.hours) {
        const hour = date.hours[idx]

        if (hour.available) {
          const start = dayjs(hour.hour).tz(TRAINING_TIMEZONE)
          const end = start.add(1, 'hour')
          scheduleEvents.push({
            display: 'background',
            start: start.format(),
            end: end.format(),
          })
        }
      }
    }

    return scheduleEvents
  })
}

onMounted(() => {
  const isTrainer = getUserRole() === Trainer

  getCalendar().then((data) => {
    const events: EventInput[] = data.data.trainings.map((obj) => {
      const start = dayjs(obj.time).tz(TRAINING_TIMEZONE)

      const end = start.add(1, 'hour')

      return {
        title: isTrainer ? obj.user : 'Training',
        start: start.format(),
        end: end.format(),
      }
    })

    getScheduleCalenderEvents().then((scheduleEvents) => options.value.events = events.concat(scheduleEvents))
  })
})
</script>

<template>
  <app-layout>
    <div class="py-5 text-center">
      <h2>Trainer's schedule</h2>
      <p class="lead">Below is an example form built entirely with Bootstrap’s form controls. Each required form
        group
        has a validation state that can be triggered by attempting to submit the form without completing it.</p>
    </div>

    <FullCalendar :options="options"/>
  </app-layout>
</template>

<style scoped>
.fc-unthemed td.fc-today {
  background: #ffffff;
}
</style>