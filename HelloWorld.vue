<template>
  <div class="hello">
  <div class="hello">
     <input
      v-model="inputData"
      type="text"
      class="form-control"
      placeholder="Enter data"
      required
    />
    <input
      v-model="inputKey"
      type="text"
      class="form-control"
      placeholder="Enter key"
      required
    />

    <input type="submit" value = "Post Data" class="btn btn-primary" @click="submitData" />
    </div>
     <input type="submit" value = "Get Data" class="btn btn-primary" @click="getData" />
     <div class="hello"> key:{{postData.key}}</div>
     <div class="hello"> cid:{{postData.cid}}</div>
     <div class="hello"> encrypteddata:{{postData.data}}</div>
     <div class="hello"> decrypteddata:{{getResData}}</div>

    
  </div>
</template>

<script>
import axios from "axios";

export default {
 
  data() {
    return { postData:"" ,getResData:""};
  },



  methods: {
    submitData() {
      console.log(this.inputData);
      if (this.inputData != "" && this.inputKey != "") {
        const body = {
          data: this.inputData,
          key: this.inputKey,
        };
        console.log("body",body)
      
        axios.post('http://localhost:1323/add', body, {
   headers: {
  "Access-Control-Allow-Origin": "*",
  "Access-Control-Allow-Methods": "GET, POST, PATCH, PUT, DELETE, OPTIONS",
  "Access-Control-Allow-Headers": "Origin, Content-Type, X-Auth-Token"
	}
      }).then(res => {
	console.log(res.data.data);
	this.postData = res.data
      }).catch(err => {
	console.log(err.response);
      });
      } else {
        alert("Both the fields are mandetory!");
      }
    },
    getData() {
            const body = {
          key: this.inputKey,
        };
    console.log(this.postData.cid)
    console.log(body)
    axios.get('http://localhost:1323/cid/'+this.postData.cid + `/${this.inputKey}`, {
   headers: {
  "Access-Control-Allow-Origin": "*",
  "Access-Control-Allow-Methods": "GET, POST, PATCH, PUT, DELETE, OPTIONS",
  "Access-Control-Allow-Headers": "Origin, Content-Type, X-Auth-Token"
	}
      }).then(res => {
      this.getResData = res.data;
      console.log(this.getResData)
    });
    

    }
    }
  };

 
</script>



