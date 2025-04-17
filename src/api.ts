import axios from 'axios'

const baseUrl = 'https://rninl-92-126-119-99.a.free.pinggy.link/api/v1'

export async function loginApi(username: string, password: string) {
  let res
  await axios
    .post(baseUrl + '/auth/login', {
      username: username,
      password: password,
    })
    .then((r) => (res = r.data))
    .catch((e) => alert(e))
  return res
}

export async function getPids() {
  let res
  await axios
    .get(baseUrl + '/pids')
    .then((r) => (res = r.data))
    .catch((e) => alert(e))
  return res
}

export async function savePidsApi(pids: any) {
  let res
  await axios
    .patch(baseUrl + '/pids', { data: pids })
    .then((r) => (res = r.data))
    .catch((e) => alert(e))
  return res
}
