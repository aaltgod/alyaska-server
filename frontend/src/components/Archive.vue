<template>
    <div class="archive">
        &nbsp;
        <button v-on:click="getFilesData()">
            Archive
        </button>
        <h3 v-if="prevPath">
           <a v-on:click="getFilesData(prevPath)"><p> ../ </p></a> 
        </h3>
        <div v-for="file in files" v-bind:key="file.path">
          <h3 v-if="file.dir">
                <a v-on:click="getFilesData(file.path)"><p> {{ file.name }}/</p></a>
            </h3>
          <h3 v-else>
            {{ file.name }} 
              <vs-button color="dark" type="gradient" v-on:click="getFile(file.path, file.name)">
              </vs-button>
          </h3>
          
        </div>
    <hr>
    </div>
</template>


<script>

import axios from 'axios'

export default {
  data: function() {
    return {
      files: [],
      prevPath: "",
    }
  },

  methods: {
    getrandomstr: function() {
      var data = {}
      axios.post("http://127.0.0.1:3000/api/random_string", data).then(result => {
        this.result = result.data['ex']
      }).catch( error =>  {
        console.error(error)
      })
    },

    getFilesData: function(filePath) {
      axios.post(
        "http://127.0.0.1:3000/api/files",
          {"path": filePath}
          )
      .then(response => {
        this.files = response.data["files"]
        this.prevPath = this.files[0]["previous_path"]
      })
      .catch(e => {
        console.error(e)
      })
    },
    
    getFile: function(filePath, fileName) {
      axios({
        method: "post",
        url: "http://127.0.0.1:3000/api/get-file",
        responseType: "arraybuffer",
        data: {
          "path": filePath,
        }
      })
      .then(response => {
        var fileURL = window.URL.createObjectURL(new Blob([response.data]))
        var fileLink = document.createElement('a')

        fileLink.href = fileURL
        fileLink.setAttribute("download", fileName)
        document.body.appendChild(fileLink)

        fileLink.click()
        })
      .catch(e => {
        console.error(e)
      })
    }
  }
}
</script>