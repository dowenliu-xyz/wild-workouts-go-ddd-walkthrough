export const serverSettings = {
    hostname: window.location.hostname,
}

export const tokenProvider = (): string => {
    return localStorage.getItem('token') ?? ''
}
