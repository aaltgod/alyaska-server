<template>
    <div class="archive">
        &nbsp;
        <button v-on:click="getFiles()">
            Archive
        </button>
        <div v-for="file in files" v-bind:key="file.Path">
          <h3 v-if="file.Dir">
                <a v-on:click="getFiles()"><p>{{ file.Name }} {{ file.Path }}</p></a>
            </h3>
          <h3 v-else>
            {{ file.Name }}
          </h3>
        </div>
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
      files: [],
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

    getFiles: function() {
      axios.get("http://127.0.0.1:3000/files", {
        headers: {
        }
      })
      .then(response => {
        this.files = response.data["files"]
      })
      .catch(e => {
        console.error(e)
      })
    }
  }
}
</script>