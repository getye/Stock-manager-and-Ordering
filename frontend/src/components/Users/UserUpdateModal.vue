<script setup>
import { defineProps, watch, watchEffect } from 'vue';
import { FwbModal } from 'flowbite-vue'
import { ref, computed, onMounted } from 'vue';
import { z } from 'zod';
const props = defineProps({
    isUpdateModal: Boolean,
    updateModal: Function,
    selectedUser: {
        type: Object,
        required: true,
  },
})

const users = ref([]);
const updatedUser = ref({
    name:'',
    email:'',
    role:'',
    phone:''
})



watch(async () => {
    try {
        updatedUser.value.name = props.selectedUser.name
        updatedUser.value.role = props.selectedUser.role
        updatedUser.value.email = props.selectedUser.email
        updatedUser.value.phone = props.selectedUser.phone
    } catch (error) {
        console.error('Error:', error)
    }
})

const handleUpdate = async ( ) => {
    /*
    const token = localStorage.getItem('token')
    if(!token){
        console.log('Error: No token')
      }
    */

    try {
        const id = props.selectedUser.id
        const response = await fetch(`http://localhost:5000/api/users/update/${id}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
                //'Authorization': `Bearer ${token}`,
            },
            body: JSON.stringify(updatedUser.value),
        });
        console.log(response)
        if (response.ok) {
            alert('Updated Successfully')
        } else {
            alert(`Error: ${response.statusText}`)
        }
    } catch (err) {
        alert('Error occurred while updating')
        console.error(err)
    }
  };

</script>

<template>
    <fwb-modal v-if="isUpdateModal" @close="updateModal" size="md"
        class="fixed inset-x-0 inset-y-0 z-50 flex items-center justify-center pl-56">
        <template #header>
            <h2>Update user role</h2>
        </template>
        <template #body>
            <div class="w-full max-w-md px-8 space-y-3 ">
                <div class="mb-4">
                    <input
                        id="userName"
                        v-model="updatedUser.name"
                        type="text"
                        placeholder="User Name"
                        class="block w-full mt-1 border-blue-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                    />
                </div>

                <div class="mb-4">
                    <input
                        id="email"
                        v-model="updatedUser.email"
                        type="email"
                        placeholder="Email"
                        class="block w-full mt-1 border-blue-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                    />
                </div>

                <div class="mb-4">
                    <input
                        id="phone"
                        v-model="updatedUser.phone"
                        type="text"
                        placeholder="Phone Number"
                        class="block w-full mt-1 border-blue-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                    />
                </div>

                <div class="mb-4">
                    <select
                        id="role"
                        v-model="updatedUser.role"
                        class="block w-full mt-1 border-blue-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm">
                        <option value="" disabled>Select a role</option>
                        <option value="Admin">Admin</option>
                        <option value="Supplier">Supplier</option>
                    </select>
                </div>
                <div class="flex justify-between space-x-4">
                    <button
                        @click="updateModal"
                        type="button"
                        class="px-4 py-1 font-medium text-white bg-red-500 rounded-lg hover:bg-red-600"
                        >
                        Close
                    </button>
                    <button
                        @click="handleUpdate"
                        class="px-5 py-1 font-medium text-white bg-blue-500 rounded-lg hover:bg-blue-600"
                        >
                        Update
                    </button>
                </div>
            </div>
        </template>
    </fwb-modal>
</template>

<style scoped>
.fade-enter-active, .fade-leave-active {
    transition: opacity 0.5s;
}
.fade-enter-from, .fade-leave-to {
    opacity: 0;
}
</style>

