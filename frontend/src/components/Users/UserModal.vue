<script setup>
import { defineProps } from 'vue';
import { FwbModal } from 'flowbite-vue'
import { ref, onMounted } from 'vue';
import { z } from 'zod';

import { toast } from 'vue3-toastify';
import 'vue3-toastify/dist/index.css';
defineProps(['isShowModal', 'showModal' ])


const newUser = ref({ name: '', email: '', phone: '', role: '', status:"Active" })
const errors = ref({})

const userSchema = z.object({
    name: z.string().min(1, 'User Name is required'),
    email: z.string().email('Invalid email address'),
    phone: z
        .string()
        .min(10, 'Phone number must be at least 10 digits')
        .max(15, "Phone number can't exceed 15 digits")
        .regex(/^\+?[0-9]+$/, 'Phone number should only contain digits and can start with +'),
    role: z.string().min(1, 'Role is required'),
})


const validateForm = () => {
  try {
        userSchema.parse(newUser.value)
        errors.value = {}
        return true;
  } catch (e) {
        if (e instanceof z.ZodError) {
        errors.value = e.errors.reduce((acc, error) => {
            acc[error.path[0]] = error.message
            return acc;
        }, {})
        }
        return false;
  }
}



const handleAddUser = async () => {
    if (validateForm()) {
        try {
            /*
            const token = localStorage.getItem('token')

            if(!token){
                console.log('Error: No token')
            } */
            const response = await fetch('http://localhost:5000/api/users/add', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    //'Authorization': `Bearer ${token}`,
                },
                body: JSON.stringify(newUser.value),
            });

            if (response.status === 201) {
                const res = await response.json()
                alert(res.message)
                newUser.value.name = ''
                newUser.value.email = ''
                newUser.value.role = ''
                newUser.value.phone = ''
            }else{
                const res = await response.json()
                alert(res.error)
                console.error('Error in registration:', res.error)
            }
        } catch (error) {
            alert('An unexpected error occurred')
            console.error('Error adding user:', error.message)
        }
    }
}

</script>

<template>

    <fwb-modal v-if="isShowModal" @close="showModal" size="md"
        class="fixed inset-x-0 inset-y-0 z-50 flex items-center justify-center pl-56">
        <template #header>
            <h2>Add a new user</h2>
        </template>
        <template #body>
            <div class="w-full max-w-md px-2 space-y-3 ">
                <div>
                    <div class="mb-4">
                        <input
                            id="userName"
                            v-model="newUser.name"
                            type="text"
                            placeholder="User Name"
                            class="block w-full mt-1 border-blue-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                        />
                    </div>

                    <div class="mb-4">
                        <input
                            id="email"
                            v-model="newUser.email"
                            type="email"
                            placeholder="Email"
                            class="block w-full mt-1 border-blue-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                        />
                    </div>

                    <div class="mb-4">
                        <input
                            id="phone"
                            v-model="newUser.phone"
                            type="text"
                            placeholder="Phone Number"
                            class="block w-full mt-1 border-blue-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                        />
                    </div>

                    <div class="mb-4">
                        <select
                            id="role"
                            v-model="newUser.role"
                            class="block w-full mt-1 border-blue-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm">
                            <option value="" disabled>Select a role</option>
                            <option value="Admin">Admin</option>
                            <option value="Supplier">Supplier</option>
                        </select>
                    </div>

                    <div class="flex justify-end space-x-4">
                        <button
                            @click="handleAddUser"                            
                            class="px-5 py-1 font-medium text-white bg-blue-500 rounded-lg hover:bg-blue-600">
                            Add
                        </button>
                    </div>
                </div>
            </div>
        </template>

    </fwb-modal>
</template>
