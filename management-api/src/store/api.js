import httpaccess from '@/util/httpaccess'
export default{
  namespaced: true,
  actions: {
    index ({ commit }) {
      return new Promise(resolve => {
        httpaccess.get(`${process.env.API_ENDPOINT}/api?offset=0&limit=10`, {})
          .then(response => {
            resolve(response.data)
          })
          .catch(e => {
            this.errors.push(e)
          })
      })
    }
  }
}
