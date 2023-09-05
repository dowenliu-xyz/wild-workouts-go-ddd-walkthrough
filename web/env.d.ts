/// <reference types="vite/client" />

interface ImportMetaEnv {
    VITE_USERS_API_BASE_PATH?: string | null
    VITE_TRAININGS_API_BASE_PATH?: string | null
    VITE_TRAINER_API_BASE_PATH?: string | null
    VITE_FIREBASE_AUTH_EMULATOR_HOST? : string | null
    VITE_FIRESTORE_EMULATOR_HOST? : string | null
}

interface ImportMeta {
    readonly env: ImportMetaEnv
}