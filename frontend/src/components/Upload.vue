<template>
    <div class="upload">
         <!-- <form
            enctype="multipart/form-data"
            action="http://127.0.0.1:3000/api/upload-file"
            method="post"
        >
            <input type="file" name="myFile"/>
            <input type="submit" value="upload-file"/>
        </form> -->
         <div class="container">
            <div class="large-12 medium-12 small-12 cell">
                <input type="file" name="files" id="files" ref="files" multiple v-on:change="handleFileUpload()"/>
                <button type="submit" v-on:click="submitFile()">Upload</button>
            </div>

            <div class="large-12 medium-12 small-12 cell">
                <button v-on:click="addFiles()">Add Files</button>
            </div>
        </div>
        <!-- <div v-if="currentFile" class="progress">
            <div
                class="progress-bar progress-bar-info progress-bar-striped"
                    role="progressbar"
                    :aria-valuenow="progress"
                    aria-valuemin="0"
                    aria-valuemax="100"
                    :style="{ width: progress + '%' }"
            >
                {{ progress }}%
            </div>
        </div>

        <label class="btn btn-default">
            <input type="file" ref="file" @change="selectFile" />
        </label>

        <button class="btn btn-success" :disabled="!selectedFiles" @click="upload">
                Upload
        </button>

        <div class="alert alert-light" role="alert">{{ message }}</div>
        <div class="card">
            <div class="card-header">List of Files</div>
            <ul class="list-group list-group-flush">
                <li
                    class="list-group-item"
                    v-for="(file, index) in fileInfos"
                    :key="index"
                >
                    <a :href="file.url">{{ file.name }}</a>
                </li>
            </ul>
        </div> -->
    </div>
</template>

<script>
import axios from 'axios'
    export default {
        data(){
            return {
                files:""
            }
        },

        methods: {
            submitFile(){
                let formData = new FormData()

                 for( var i = 0; i < this.files.length; i++ ){
                    let file = this.files[i]

                    formData.append('files[' + i + ']', file)
                    }

                // formData.append("file", this.file)

                axios.post(process.env.VUE_APP_API_URL+"/api/upload-file",
                formData,
                {
                    headers: {
                        'Content-Type': 'multipart/form-data',
                    }

                }).then(function(){
                    console.log("SUC")
                })
                .catch(function(){
                    console.log("FAIL")
                })
            },

            handleFileUpload(){
                this.files = this.$refs.files.files
                console.log(this.files)
            }
        }
    }
</script>

