import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    redirect: '/login',
  },
  {
    path: '/login',
    name: 'login',
    component: () => import('../views/Login.vue'),
  },
  {
    path: '/trainings',
    name: 'trainingsList',
    component: () => import('../views/TrainingsList.vue'),
  },
  {
    path: '/calendar',
    name: 'calendar',
    component: () => import('../views/CalendarView.vue'),

  },
  {
    path: '/trainings/schedule',
    name: 'scheduleTraining',
    component: () => import('../views/ScheduleTraining.vue')
  },
  {
    path: '/trainings/reschedule/:trainingID',
    name: 'rescheduleTraining',
    component: () => import('../views/ScheduleTraining.vue'),
  },
  {
    path: '/trainings/propose-new-date/:trainingID',
    name: 'proposeNewDate',
    component: () => import('../views/ScheduleTraining.vue'),
    props: {isPropose: true},
  },
  {
    path: '/trainer/set-schedule',
    name: 'setSchedule',
    component: () => import('../views/SetSchedule.vue'),
  },
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: routes
})

export default router
