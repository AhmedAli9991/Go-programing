<script>
import axios from "axios"

  export default {
    async mounted() {
      const url = "http://localhost:5000/user"
      await axios.get(url, {
        withCredentials: true
      }).then((response) => {
        if (response.status === 200) {
          this.user = response.data.name
          this.email = response.data.email
        }
      })
    },
    data() {
      return {
        user: "",
        email: ""
      }
    }
  }
      

</script>

<template>
  <main>
    <h1>Hello, {{ user }} </h1>
    <p>Your email is {{ email }} </p>
  </main>
</template>
