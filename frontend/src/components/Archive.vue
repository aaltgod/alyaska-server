<template>
    <div class="archive">
        &nbsp;
        <button v-on:click="getFiles()">
            Archive
        </button>
        <div v-for="file in files" v-bind:key="file.path">
          <h3 v-if="file.dir">
                <a v-on:click="getFiles(file.path)"><p> {{ file.name }}/</p></a>
            </h3>
          <h3 v-else>
            {{ file.name }}
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

    getFiles: function(filePath) {

      console.log(filePath)

      axios.post(
        "http://127.0.0.1:3000/api/files",
          {"path": filePath}
          )
      .then(response => {
        this.files = response.data["files"]
        for (var file in this.files) {
          console.log(this.files[file])
        }
      })
      .catch(e => {
        console.error(e)
      })
    }
  }
}
</script>