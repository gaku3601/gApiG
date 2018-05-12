import axios from 'axios'

export default{
  request (method, url, params) {
    var promise = null
    var instance = axios.create({
      timeout: 10 * 60 * 1000,
      headers: {'Authorization': 'Token ' + localStorage.getItem('token')}
    })

    if (method === 'get') {
      promise = instance.get(url, { params: params })
    } else if (method === 'delete') {
      promise = instance.delete(url, { params: params })
    } else if (method === 'post') {
      promise = instance.post(url, params)
    } else if (method === 'patch') {
      promise = instance.patch(url, params)
    }
    promise.catch(function () {
      return alert('エラーが発生しました')
    })
    return promise
  },

  get (url, params) {
    return this.request('get', url, params)
  },

  delete (url, params) {
    return this.request('delete', url, params)
  },

  post (url, params) {
    return this.request('post', url, params)
  },

  patch (url, params) {
    return this.request('patch', url, params)
  }
}
