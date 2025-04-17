import axios from 'axios'

const baseUrl = 'http://192.168.62.236:8000/api/v1'

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

export async function savePidsApi(pids) {
  let res
  await axios
    .patch(baseUrl + '/pids', { data: pids })
    .then((r) => (res = r.data))
    .catch((e) => alert(e))
  return res
}
