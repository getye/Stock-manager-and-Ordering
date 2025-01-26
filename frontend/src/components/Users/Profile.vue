<script setup>
import { ref, onMounted, watch } from 'vue';
import { z } from 'zod';
import GuestLayout from '@/Layouts/GuestLayout.vue';
import { Head } from '@inertiajs/vue3';


const profile = ref([])


watch( async() => {
    try {
        const token = localStorage.getItem('token')

        if(!token){
            console.log('Error: No token')
        }
        const response = await fetch(`api/users/profile`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${token}`,
            }
        });

        const data = await response.json()
        profile.value = data.data

    } catch (err) {
        console.log(err)
    }
  });



</script>

<template>
    <div>
        <table class="max-w-4xl">
            <tr>
                <td class="p-3 font-semibold border-2 border-gray-200">Name </td>
                <td class="p-3 font-semibold border-2 border-gray-200">{{ profile[0].user_name }}</td>
            </tr>
            <tr>
                <td class="p-3 font-semibold border-2 border-gray-200">Email</td>
                <td class="p-3 font-semibold border-2 border-gray-200"> {{ profile[0].user_email }}</td>
            </tr>
            <tr>
                <td class="p-3 font-semibold border-2 border-gray-200">Phone</td>
                <td class="p-3 font-semibold border-2 border-gray-200"> {{ profile[0].user_phone }}</td>
            </tr>
            <tr>
                <td class="p-3 font-semibold border-2 border-gray-200 ">Role </td>
                <td class="p-3 font-semibold border-2 border-gray-200"> {{ profile[0].role.role_name }}</td>
            </tr>

            <tr>
                <td class="p-3 font-semibold border-2 border-gray-200 ">Permisions</td>
                <td class="p-3 font-semibold border-2 border-gray-200">{{ JSON.parse(profile[0].role.permissions).join(', ')}}</td>
            </tr>
        </table>
    </div>
</template>


