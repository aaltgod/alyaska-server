<template>
    <div class="upload">
        <div class="container">
            <div class="large-12 medium-12 small-12 cell">
                <label>Files
                    <input type="file" id="files" ref="files" multiple v-on:change="handleFileUpload()"/>
                </label>
            </div>
            <div class="large-12 medium-12 small-12 cell">
                <div v-for="(file, key) in files" class="file-listing" v-bind:key="file">{{ file.name }} <span class="remove-file" v-on:click="removeFile( key )">Remove</span></div>
            </div>
            <br>
            <div class="large-12 medium-12 small-12 cell">
                <button v-on:click="addFile()">Add Files</button>
            </div>
            
            <br>
            <div class="large-12 medium-12 small-12 cell">
                <button v-on:click="submitFile()">Submit</button>
            </div>
            <h3><a v-bind:href="'/folder/'+folderName">{{ folderName }}</a></h3>  
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
            }
        },

        methods: {
            addFile(){
                this.$refs.files.click()
            },

            submitFile(){
                let formData = new FormData()

                if (this.files.length == 0) 
                    return

                for( var i = 0; i < this.files.length; i++ ){
                    let file = this.files[i]

                    formData.append(i, file)
                    }

                axios.post(process.env.VUE_APP_API_URL+"/api/upload-file",
                formData,
                {
                    headers: {
                        'Content-Type': 'multipart/form-data',
                    }
                }).then(result =>{
                    this.folderName = result.data["folderName"]
                }).catch(e =>{
                    console.error(e)
                })
                
            },

            handleFileUpload(){
                let uploadedFiles = this.$refs.files.files
                
                for( var i=0; i < uploadedFiles.length; i++ ){
                    this.files.push( uploadedFiles[i] )
                }
            },

            removeFile( key ){
                this.files.splice( key, 1 )
            }
        }
    }
</script>

