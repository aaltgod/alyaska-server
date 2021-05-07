<template>
    <div class="archive">
        &nbsp;
        <button v-on:click="getFilesData()">
          Get Archive
        </button>
        <h3 v-if="prevPath">
           <span style="cursor: pointer" v-on:click="getFilesData(prevPath)"><p> ../ </p></span> 
        </h3>
        <div v-for="file in files" v-bind:key="file.path">
          <h3 v-if="file.dir">
                <span style="cursor: pointer" v-on:click="getFilesData(file.path)">{{ file.name }}/</span>
            </h3>
          <h3 v-else>
            {{ file.name }}
             <button v-on:click="getFile(file.path, file.name)">
               <i class="fas fa-file-download"></i>
            </button> 
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
    getRandomStr: function() {
      var data = {}
      axios.post(process.env.VUE_APP_API_URL+"/api/random-string", data).then(result => {
        this.result = result.data['ex']
      }).catch( error =>  {
        console.error(error)
      })
    },

    getFilesData: function(filePath) {
      axios.post(
        process.env.VUE_APP_API_URL+"/api/files",
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
        url: process.env.VUE_APP_API_URL+" /api/get-file",
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