<template>
    <div class="folder">
        <h1> {{ $route.params.folderName }} </h1>
        <div v-for="file in files" v-bind:key="file.name">
          <h3>
            {{ file.name }}
             <button v-on:click="getFile(file.path, file.name)">
               <i class="fas fa-file-download"></i>
            </button> 
          </h3>
        </div>
    </div>
</template>

<script>
import axios from 'axios'

    export default{
        data: function() {
            return {
                files: []
            }
        },

        mounted: function() {
            this.getFolder()
        },

        methods: {
            getFolder: function() {
                axios.get(
                    process.env.VUE_APP_API_URL+"/api/folder/" + this.$route.params.folderName,
                ).then(result => {
                    this.files = result.data["files"]
                }).catch( error => {
                    console.error(error)
                })
            },
            getFile: function(filePath, fileName) {
                axios({
                    method: "post",
                    url: process.env.VUE_APP_API_URL+"/api/get-file",
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
            },
    }


</script>