<template>
    <div class="archive">
        <h1>Hi there</h1>
        &nbsp;
        <p v-on:click="getEx()">
            Archive
        </p>
        <h1>{{ result }}</h1>

        <!-- <div>
          <b-button variant="primary" v-on:click="getEx()">Random {{ result }}</b-button>  
        </div>

        &nbsp; -->
    <hr>
    </div>
</template>



<script>
import axios from 'axios'

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
        this.result = result.data['ex']
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