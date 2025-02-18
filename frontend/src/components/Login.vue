<script setup>
import { ref } from 'vue';
import ForgotPassword from './ForgotPassword.vue';


const user = ref({
    email: '',
    password: '',
    remember: false,
})

const isForgotPasswordModal = ref(false)

const forgotPasswordModal = () =>{
    isForgotPasswordModal.value = true
}
const forgotPasswordModalClose = () =>{
    isForgotPasswordModal.value = false
}


const handleLogin = async () => {
    try {
        const response = await fetch('http://localhost:5000/api/users/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(user),
        })

        const responseData = await response.json()

        if (responseData.redirect) {
            window.location.href = responseData.redirect
        }

        if (response.ok) {

            localStorage.setItem('userId', responseData.user.id)
            localStorage.setItem('userRole', responseData.user.role)

        } else {
            console.log(responseData.error)
            alert('User not found')
        }
    } catch (error) {
        alert('An unexpected error occurred')
        console.error(error)
    }
}


</script>

<template>
    <div class="flex justify-center mt-20">
        <div class="max-w-md p-8 bg-gray-100 rounded-xl">
            <div>
                <input
                    id="email"
                    type="email"
                    v-model="user.email"
                    placeholder="User name"
                    required
                    class="block p-2 mt-1 border-2 border-blue-300 rounded-md min-w-96" />
            </div>

            <div class="mt-4">
                <input
                    id="password"
                    type="password"
                    v-model="user.password"
                    placeholder="Password"
                    required
                    class="block w-full p-2 mt-1 border-2 border-blue-300 rounded-md" />
            </div>

            <div class="block mt-4">
                <label class="flex items-center">
                    <input type="checkbox" name="remember" v-model="user.remember" class="text-blue-300 border-2 border-blue-300" />
                    <span class="text-gray-600 text-md ms-2">Remember me</span>
                </label>
            </div>

            <div class="flex items-center justify-between mt-4">
                <button @click="forgotPasswordModal" class="text-blue-500 text-md">
                    Forgot your password?
                </button>

                <button @click="handleLogin" class="px-3 py-1 font-semibold text-white bg-blue-400 rounded-xl">
                    Log in
                </button>
            </div>
        </div>

        <ForgotPassword
            :isForgotPasswordModal="isForgotPasswordModal"
            :forgotPasswordModalClose="forgotPasswordModalClose"
        />
    </div>

</template>
