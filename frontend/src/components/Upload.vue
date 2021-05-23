<template>
    <div class="upload">
        <div class="container">
            <div class="large-12 medium-12 small-12 cell">
                <label>Files
                    <input type="file" id="files" ref="files" multiple v-on:change="handleFileUpload()"/>
                </label>
            </div>
            <div class="large-12 medium-12 small-12 cell">
                <div v-for="(file, key) in files" class="file-listing" v-bind:key="file">
                    {{ file.name }} 
                    <i class="fa fa-window-close" style="cursor: pointer" aria-hidden="true" v-on:click="removeFile( key )"></i>
                </div>
            </div>
            <br>
            
            <br>
            <div class="large-12 medium-12 small-12 cell">
                <button v-on:click="submitFile()">Submit</button>
            </div>
            <h3 v-if="loading">
                <i class="fa fa-spinner fa-pulse fa-3x fa-fw"></i>
                <span class="sr-only">Loading...</span>
            </h3>
            <h3><a v-bind:href="'/folder/'+folderName">{{ folderName }}</a></h3>
            <h3><a v-if="filesSizeError">Files size > 250Mb</a></h3> 
        </div>
    </div>
</template>

<script>
import axios from 'axios'

    export default {
        data(){
            return {
                files:[],
                folderName: "",
                filesSizeError: false,
                loading: false,
            }
        },

        methods: {
            submitFile(){
                let formData = new FormData()

                if (this.files.length == 0) {
                    return
                }

                var filesSize = 0 

                for( var i = 0; i < this.files.length; i++ ){
                    let file = this.files[i]
                    
                    filesSize += file.size

                    formData.append(i, file)
                }

                this.filesSizeError = false

                if (filesSize > 262144000) {
                    this.filesSizeError = true
                    return
                }

                this.loading = true

                axios.post(process.env.VUE_APP_API_URL+"/api/upload-file",
                formData,
                {
                    headers: {
                        'Content-Type': 'multipart/form-data',
                    }
                }).then(result =>{
                    this.folderName = result.data["folderName"]

                    this.loading = false
                }).catch(e =>{
                    console.error(e)

                    this.loading = false
                })
            },

            handleFileUpload(){
                let uploadedFiles = this.$refs.files.files
                for( var i = 0; i < uploadedFiles.length; i++ ){
                    this.files.push( uploadedFiles[i] )
                }
                
            },

            removeFile( key ){
                this.files.splice( key, 1 )
            }
        }
    }
</script>

