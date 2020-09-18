import Vue from 'vue';
import axios from 'axios';

const BASE_URI = 'https://blog.codin.ro:1333';
const client = axios.create({
  baseURL: BASE_URI,
  json: true
});



const APIClient =  {
    login(username, password) {
      return this.perform('post', '/login', [username, data]);
    },
  
    async login_a (username, password) {
        axios.post(BASE_URI )
    },

    async perform (method, resource, data) {
      let accessToken = await Vue.prototype.$auth.getAccessToken()
      return client({
        method,
        url: resource,
        data,
        headers: {
          Authorization: `Bearer ${accessToken}`
        }
      }).then(req => {
        return req.data
      })
    }
  }
  
  export default APIClient;