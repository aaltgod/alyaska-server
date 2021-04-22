<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <p>
      The Frame will be here
    </p>
    <h3>FILES</h3>
    <ul>
      <li><a href="/files" rel="noopener">Files</a></li>
      <li><a href="/" target="_blank" rel="noopener">...</a></li>
      <li><a rel="noopener">{{ result }}</a></li>
    </ul>
    <h3>INFO</h3>

    &nbsp;
    <div>
      <b-button variant="primary" v-on:click="getEx()">Random string</b-button>  
    </div>

    &nbsp;
    <hr>

  </div>
  
</template>

<script>
import axios from 'axios';
// import Vue from 'vue'
// import VeeValidate from 'vee-validate'

/* eslint-disable */
// Vue.use(VeeValidate)

export default {
  props: {
    msg: String
  },

  data: function() {
    return {
      result: ""
    }
  },

  methods: {
    getrandomstr: function() {
      var data = {}
      axios.post("http://127.0.0.1:3000/api/random_string", data).then(result => {
        this.result = result.data['string']
      }).catch( error =>  {
        console.error(error);
      });
    },

    getEx: function() {
      axios.get("http://127.0.0.1:3000/api/ex", {
        headers: {
        }
      })
      .then(response => {
        this.result = response.data["ex"]
      })
      .catch(e => {
        console.error(e)
      })
    }
  }
}

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
