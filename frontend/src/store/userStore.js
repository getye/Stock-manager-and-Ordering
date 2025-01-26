import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export const useUserStore = defineStore('user', () => {
    const storedUsers = ref([])
    const searchTerm = ref('');

        function setStoredUsers(users){
            storedUsers.value = users
        }

        function setSearchTerm(term) {
            searchTerm.value = term;
        }

        const filteredUsers = computed(() => {
            const term = searchTerm.value.toLowerCase();
            return storedUsers.value.filter(
            (user) =>
                user.name?.toLowerCase().includes(term) ||
                user.email?.toLowerCase().includes(term) ||
                user.role?.toLowerCase().includes(term)||
                user.status?.toLowerCase().includes(term)
            )
        })

    return {
        storedUsers, setStoredUsers,
        searchTerm, setSearchTerm,
        filteredUsers,
    }
  })
