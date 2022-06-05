<template>
  <div>
      <form @submit.prevent="">
        <div>
          <input 
            type="email" 
            name="email" 
            id="email" 
            placeholder="email"
            v-model="data.email">
          <input 
            type="password" 
            name="password" 
            id="password" 
            placeholder="password"
            v-model="data.password">
        </div>
        <div>
          <button @click="submit">Login</button>
        </div>
      </form>
  </div>
</template>

<script>
import { reactive } from "vue"
import { useRouter } from "vue-router"

export default {

  setup() {
    const data = reactive({
      email: "",
      password: "",
    })
    const router = useRouter()
    return {
      data,
      router,
    }
  },
  methods: {
    async submit() {
      
      const url = "http://localhost:5000/auth/login"
      const res = await fetch(url,  {
        method: "POST",
        headers: {"Content-Type": "application/json"},
        credentials: "include",
        body: JSON.stringify(this.data)
      })
      const status = res.status
      if (status == 200) {
        
        this.router.push("/")
      } else {
        console.log("no auth")
      }
    }
  }
}
</script>

<style>

</style>