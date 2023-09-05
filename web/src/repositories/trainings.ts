import {
    Configuration as TrainerApiConfiguration,
    DefaultApi as TrainerApi,
    type HourUpdate,
    type ModelDate,
} from '@/repositories/clients/trainer'
import {
    Configuration as TrainingsApiConfiguration,
    DefaultApi as TrainingsApi,
    type PostTraining,
    type Trainings,
} from '@/repositories/clients/trainings'
import {serverSettings, tokenProvider} from './config'
import type {AxiosResponse} from 'axios'
import {TRAINING_TIMEZONE, UTC_FORMAT} from '@/date'
import dayjs from 'dayjs'
import utc from 'dayjs/plugin/utc'
import timezone from 'dayjs/plugin/timezone'

dayjs.extend(utc)
dayjs.extend(timezone)

const getTrainerApiBasePath = (): string => {
    return import.meta.env.VITE_TRAINER_API_BASE_PATH ?? (serverSettings.hostname + '/api')
}

const trainerApi = new TrainerApi(new TrainerApiConfiguration({
    basePath: getTrainerApiBasePath(),
    accessToken: tokenProvider,
}))

export const getSchedule = async (dateFrom: string, dateTo: string): Promise<Array<ModelDate>> => {
    return trainerApi.getTrainerAvailableHours(dateFrom, dateTo)
        .then((response) => response.data)
}

export const getAvailableDates = async (): Promise<Array<ModelDate>> => {
    const from = dayjs().utc().set('hour', 0).set('minute', 0).set('second', 0).set('millisecond', 0)
    const to = from.add(3, 'week')

    return getSchedule(from.format(UTC_FORMAT), to.format(UTC_FORMAT))
        .then((data) => data.filter(day => day.hasFreeHours))
}

export const setHourAvailability = async (updates: string[][], availability: boolean): Promise<void> => {
    const hourUpdates: string[] = []

    updates.forEach((val) => {
        hourUpdates.push(val[0] + 'T' + val[1] + ':00Z')
    })

    const hourUpdate: HourUpdate = {
        hours: hourUpdates,
    }

    if (availability) {
        return trainerApi.makeHourAvailable(hourUpdate).then(() => {
        })
    } else {
        return trainerApi.makeHourUnavailable(hourUpdate).then(() => {
        })
    }
}

const getTrainingsApiBasePath = (): string => {
    return import.meta.env.VITE_TRAININGS_API_BASE_PATH ?? (serverSettings.hostname + '/api')
}

const trainingsApi = new TrainingsApi(new TrainingsApiConfiguration({
    basePath: getTrainingsApiBasePath(),
    accessToken: tokenProvider,
}))

export const getCalendar = async (): Promise<AxiosResponse<Trainings>> => {
    return trainingsApi.getTrainings()
}

export interface Period {
    from: string
    to: string
}

export const getPeriods = (): Period[] => {
    const periods: Period[] = []

    for (let week = 0; week <= 2; week++) {
        const from = dayjs().utc().add(week * 8, 'day')
        const to = from.add(7, 'day')

        periods.push({from: from.tz(TRAINING_TIMEZONE).format('YYYY-MM-DD'), to: to.tz(TRAINING_TIMEZONE).format('YYYY-MM-DD')})
    }

    return periods
}

export const scheduleTraining = async (notes: string, date: string, hour: string): Promise<void> => {
    const req: PostTraining = {
        notes,
        time: date + 'T' + hour + ':00Z',
    }

    return trainingsApi.createTraining(req).then(() => {
    })
}

export const rescheduleTraining = async (trainingUUID: string, notes: string, date: string, hour: string, isPropose: boolean): Promise<void> => {
    const req: PostTraining = {
        notes: notes,
        time: date + 'T' + hour + ':00Z',
    }

    if (isPropose) {
        return trainingsApi.requestRescheduleTraining(trainingUUID, req)
            .then(() => {
            })
    }
    return trainingsApi.rescheduleTraining(trainingUUID, req)
        .then((() => {
        }))
}

export const cancelTraining = async (uuid: string): Promise<void> => {
    return trainingsApi.cancelTraining(uuid).then(() => {
    })
}

export const approveReschedule = async (uuid: string): Promise<void> => {
    return trainingsApi.approveRescheduleTraining(uuid).then(() => {
    })
}

export const rejectReschedule = async (uuid: string): Promise<void> => {
    return trainingsApi.rejectRescheduleTraining(uuid)
        .then(() => {
        })
}
