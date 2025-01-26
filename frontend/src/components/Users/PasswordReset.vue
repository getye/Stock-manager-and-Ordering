<script setup>
import { ref, onMounted } from 'vue';
import { z } from 'zod';
import GuestLayout from '@/Layouts/GuestLayout.vue';
import { Head } from '@inertiajs/vue3';


const password = ref({
    currentPassword: '',
    newPassword: '',
    confirmPassword: ''
});
const errors = ref({});

const currentType = ref('password')
const newType = ref('password')
const confirmType = ref('password')

const passwordResetSchema = z.object({
    currentPassword: z
        .string()
        .min(6, { message: "Current password is required" }),
    newPassword: z
        .string()
        .min(8, { message: "New password must be at least 8 characters long" })
        .regex(/[A-Z]/, { message: "New password must contain at least one uppercase letter" })
        .regex(/[a-z]/, { message: "New password must contain at least one lowercase letter" })
        .regex(/[0-9]/, { message: "New password must contain at least one digit" })
        .regex(/[@$!%*?&#]/, { message: "New password must contain at least one special character" }),
    confirmPassword: z
        .string()
        .min(8, { message: "Confirm password must be at least 8 characters long" })})
        .refine((data) => data.newPassword === data.confirmPassword, {
            message: "Passwords do not match",
            path: ["confirmPassword"],
    })


const validateForm = () => {
  try {
    passwordResetSchema.parse(password.value)
        errors.value = {};
        return true;
  } catch (e) {
        if (e instanceof z.ZodError) {
        errors.value = e.errors.reduce((acc, error) => {
            acc[error.path[0]] = error.message
            return acc
        }, {});
        }
        return false;
  }
};


const handleResetPassword = async () => {
    if (validateForm()) {
        try {
            const token = localStorage.getItem('token')

            if(!token){
                console.log('Error: No token')
            }
            const userPass = {
                currentPassword: password.value.currentPassword,
                newPassword: password.value.newPassword
            }
            const response = await fetch(`/api/password/update`, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${token}`,
                },
                body: JSON.stringify(userPass),
            });

            if (response.status === 200) {
                alert('Password Successfully Updated');
                password.value = {
                    currentPassword: '',
                    newPassword: '',
                    confirmPassword: ''
                };
            }else{
            const res = await response.json()
            alert(res.message)
            console.error('Error:', res.error)
            }
        } catch (error) {
            alert('An unexpected error occurred')
            console.error('Error:', error)
        }
    }
}

const showCurrentPassword = ()=>{
    currentType.value = 'text'
}

const showNewPassword = ()=>{
    newType.value = 'text'
}
const showConfirmPassword = ()=>{
    confirmType.value = 'text'
}

</script>

<template>
    <div>
        <div class="w-full px-8 space-y-3 ">
            <div class="relative ">
                <input
                    v-model="password.currentPassword"
                    :type="currentType"
                    placeholder="Current Password"
                    class="block p-3 mt-1 border-gray-300 rounded-lg shadow-sm w-96 focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                    />
                <div class="absolute inset-y-0 right-0 flex items-center pr-4 text-white">
                    <button @click="showCurrentPassword">
                        <font-awesome-icon icon="eye" class="text-gray-500"/>
                    </button>
                </div>
            </div>
            <p v-if="errors.currentPassword" class="text-sm text-red-500">{{ errors.currentPassword }}</p>


            <div class="relative ">
                <input
                    v-model="password.newPassword"
                    :type="newType"
                    placeholder="New Password"
                    class="block w-full p-3 mt-1 border-gray-300 rounded-lg shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                    />
                <div class="absolute inset-y-0 right-0 flex items-center pr-4 text-white">
                    <button @click="showNewPassword">
                        <font-awesome-icon icon="eye" class="text-gray-500"/>
                    </button>
                </div>
            </div>
            <p v-if="errors.newPassword" class="text-sm text-red-500">{{ errors.newPassword }}</p>


            <div class="relative ">
                <input
                    v-model="password.confirmPassword"
                    :type="confirmType"
                    placeholder="Confirm Password"
                    class="block w-full p-3 mt-1 border-gray-300 rounded-lg shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                    />
                <div class="absolute inset-y-0 right-0 flex items-center pr-4 text-white">
                    <button @click="showConfirmPassword">
                        <font-awesome-icon icon="eye" class="text-gray-500"/>
                    </button>
                </div>
            </div>
            <p v-if="errors.confirmPassword" class="text-sm text-red-500">{{ errors.confirmPassword }}</p>


            <div class="flex justify-end space-x-4">
                <button
                    @click="handleResetPassword"
                    class="px-5 py-1 font-medium text-white bg-blue-500 rounded-lg hover:bg-blue-600"
                    >
                    Update
                </button>
            </div>
        </div>
    </div>
</template>

<style scoped>
.fade-enter-active, .fade-leave-active {
    transition: opacity 0.5s;
}
.fade-enter-from, .fade-leave-to {
    opacity: 0;
}
</style>

