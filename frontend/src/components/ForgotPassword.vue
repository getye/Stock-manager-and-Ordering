<script setup>
import { defineProps, watchEffect } from 'vue';
import { FwbModal } from 'flowbite-vue'
import { ref, computed, onMounted } from 'vue';
import { z } from 'zod';
const props = defineProps({
    isForgotPasswordModal: Boolean,
    forgotPasswordModalClose: Function,
})

const user = ref({
    email:''
})

const closeModal = () => {
  props.forgotPasswordModalClose();
};

const forgotPassword = async () => {
    try {
        const response = await fetch('/api/users/forgot-password', {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(user.value),
        });

        const res = await response.json();

        if (response.ok) {
            alert("OTP has sent with your email")
        } else {
            alert(res.message);
        }
    } catch (error) {
        alert('An unexpected error occurred');
        console.error(error);
    }
}

</script>

<template>
    <fwb-modal v-if="props.isForgotPasswordModal" @close="closeModal" size="md" class="fixed inset-x-0 inset-y-0 z-50 flex items-center justify-center">
        <template #header>
            <h2>Forgot Password</h2>
        </template>
        <template #body>
            <div class="w-full max-w-md space-y-3">
                <div class="py-5 space-y-3">
                    <input
                        type="email"
                        v-model="user.email"
                        placeholder="Recovery Email"
                        class="w-full p-2 border-2 border-blue-200 rounded-lg"
                    />
                </div>
                <div class="flex justify-end pb-5 pr-5">
                    <button
                        @click="forgotPassword"
                        class="px-5 py-1 font-medium text-white bg-blue-500 rounded-lg hover:bg-blue-600"
                        >
                        Recover
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

