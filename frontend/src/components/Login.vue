<template>


    <v-content>
      <v-container fluid fill-height>
        <v-layout align-center justify-center>
          <v-flex xs12 sm8 md4>

  <v-form
    ref="form"

  >
    <v-text-field
      v-model="username"
      :counter="10"
      :rules="[rules.required]"
      label="Username"
      required
    ></v-text-field>

    <v-text-field
      v-model="password"
      :rules="[rules.required]"
      :type="'password'"
      label="Password"
      required
    ></v-text-field>

    <v-btn
      color="blue"
      @click="login"
    >
      Login
    </v-btn>
  </v-form>

            </v-flex>
        </v-layout>
      </v-container>
    </v-content>


</template>
<script>
import axios from 'axios'
import Vuex from 'vuex'
// import store from 'store'
import { mapActions, mapState } from 'vuex'


export default {
    data() {
        return {
            rules: {
            required: value => !!value || 'Required.',
            },
            username: '',
            password: '',

        }
    },
    computed: {
  ...mapState([
    'loggingIn',
    'loginError'
  ])
    },
    methods: {
      ...mapActions([
        'doLogin'
      ]),
        login (){
            let currentObj = this;
            let formData = new FormData();
            formData.append('username', this.username);
            formData.append('password', this.password);
            this.doLogin(formData);
            // axios.post('http://localhost:1333/auth/login', formData).then(function (response) {
            //     currentObj.output = response.data;
            //     console.log(response.data)
            //     localStorage.setItem('token', JSON.stringify(response.data.token));
                
            // }).catch(function (error) {
            //     currentObj.output = error;
            // })
        }
    }
}
</script>