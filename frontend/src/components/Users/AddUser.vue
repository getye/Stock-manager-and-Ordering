
<script setup>
import { onMounted, ref, watch, watchEffect } from 'vue'
import UserModal from './UserModal.vue';
import UserUpdateModal from './UserUpdateModal.vue';
import { useUserStore } from '../../store/userStore';
import { toast } from 'vue3-toastify';
import 'vue3-toastify/dist/index.css';

const userStore = useUserStore()
const isShowModal = ref(false)
const isUpdateModal = ref(false)
const users = ref({})
const selectedUser = ref({})

const role = localStorage.getItem('userRole')
function showModal () {
    isShowModal.value = !isShowModal.value
}


function updateModal (user) {
    isUpdateModal.value = true
    selectedUser.value = { ...user }
}

function closeUpdateModal () {
    isUpdateModal.value = false
}


onMounted( async() => {
    //if(role === 'Admin'){
        try {
            /*
            const token = localStorage.getItem('token')

            if(!token){
                console.log('Error: No token')
            }
                */
            const response = await fetch('http://localhost:5000/api/users/view', {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                    //'Authorization': `Bearer ${token}`,
                }
            });

            const data = await response.json()
            console.log(data)
            users.value = data.users
            userStore.setStoredUsers(data.users)

        } catch (err) {
            console.log(err)
        }
    //}else toast("You have no permission")
  });



const handleUpdateStatus = async (id, user_status) => {

    /*
    const token = localStorage.getItem('token')

    if(!token){
        console.log('Error: No token')
    }
    */

    try {
        const status = (user_status == "Active")? "Inactive": "Active"
        
        const response = await fetch(`http://localhost:5000/api/users/update/status/${id}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
                //'Authorization': `Bearer ${token}`,
            },
            body: JSON.stringify({ status: status }),
        })
        const res = await response.json()
        if (response.ok) {
            alert('User Status Updated Successfully')
        } else {
            alert(`Error:`, res.error)
        }
    } catch (err) {
        alert('Error occurred while updating user status')
        console.error(err);
    }
  }

const handleDeleteUser = async (id, name) => {
    /*
    const token = localStorage.getItem('token')

    if(!token){
        console.log('Error: No token')
    }
    */
    if (!window.confirm(`Are you sure you want to delete : ${name}?`)) {
        return;
    }

    try {
        const response = await fetch(`http://localhost:5000/api/users/delete/${id}`, {
            method: 'DELETE',
            headers: {
                'Content-Type': 'application/json',
                //'Authorization': `Bearer ${token}`,
            },
        })
        console.log(response)
        if (response.ok) {
            alert('User Deleted Successfully')
        } else {
            alert(`Error: ${response.statusText}`)
        }
    } catch (err) {
        alert('Error occurred while deleting user', 'Error')
        console.error(err)
    }
};

watch(()=>{
    if(userStore.filteredUsers.length > 0){
        users.value = userStore.filteredUsers
    }
})

</script>

<template>
    <div>
        <div class="flex justify-center">
            <div  class="flex flex-col items-center w-full max-w-3xl mt-16 space-y-4">
                <div class="flex items-center justify-start w-full max-w-xl space-x-4">
                    <button
                        @click="showModal"
                        class="inline-flex gap-2 px-2 py-1 font-bold text-white bg-blue-400 rounded-lg"
                        >
                        <font-awesome-icon
                            :icon="['fas', 'user-plus']"
                            class="pt-1"
                            />
                        Add User
                    </button>
                    <input
                        type="text"
                        v-model="userStore.searchTerm"
                        placeholder="Search users ..."
                        class="py-1 pl-4 border-2 border-blue-500 w-80 rounded-2xl"
                    />
                </div>

                <div v-if="users.length > 0" class="w-full">
                    <table class="table w-full">
                        <thead class="font-semibold text-center text-white bg-blue-300 ">
                            <tr>
                                <th class="px-4 py-1 border border-gray-300">User Name</th>
                                <th class="px-4 py-1 border border-gray-300">User Role</th>
                                <th class="px-4 py-1 border border-gray-300">Email</th>
                                <th class="px-4 py-1 border border-gray-300">Phone</th>
                                <th class="px-4 py-1 border border-gray-300">Status</th>
                                <th class="px-4 py-1 border border-gray-300">Actions</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr v-for="user in users" :key="user.id">
                                <td class="px-4 py-2 border border-gray-300">{{ user.name }}</td>
                                <td class="px-4 py-2 border border-gray-300">{{ user.role }}</td>
                                <td class="px-4 py-2 border border-gray-300">{{ user.email }}</td>
                                <td class="px-4 py-2 border border-gray-300">{{ user.phone }}</td>
                                <td v-if="user.status =='Active'"
                                    class="px-4 py-2 border border-gray-300" >
                                    <button
                                        @click="handleUpdateStatus(user.id, user.status)"
                                        class="inline-flex px-2 text-sm text-blue-800 bg-blue-300 rounded-3xl ">
                                        {{ user.status }}
                                        <font-awesome-icon icon="toggle-on" class="px-1 pt-1" />
                                    </button>
                                </td>
                                <td v-else
                                    class="px-4 py-2 border border-gray-300" >
                                    <button
                                        @click="handleUpdateStatus(user.id, user.status)"
                                        class="px-2 text-sm text-red-400 bg-blue-200 rounded-3xl ">
                                        {{ user.status }}
                                        <font-awesome-icon
                                            class="text-red-300"
                                            :icon="['fas', 'toggle-off']" />
                                    </button>
                                </td>
                                <td class="flex gap-2 p-2 px-4 py-2 border border-gray-300 justify-evenly">
                                    <button
                                        @click="updateModal(user)"
                                        class="text-blue-500">
                                        <font-awesome-icon :icon="['fas', 'pen-to-square']" />
                                    </button>
                                    <button
                                        @click="handleDeleteUser(user.id, user.name)"
                                        class="text-red-500">
                                        <font-awesome-icon :icon="['fas', 'trash']" />
                                    </button>

                                </td>
                            </tr>
                        </tbody>
                    </table>
                </div>
                <div v-else>
                    No record found
                </div>

            </div>

        </div>

        <UserModal
            :isShowModal="isShowModal"
            :showModal="showModal"
            />

        <UserUpdateModal
            :isUpdateModal="isUpdateModal"
            :updateModal="closeUpdateModal"
            :selectedUser = "selectedUser"
            />

    </div>
</template>

