<script setup>
import { defineProps, watchEffect } from 'vue';
import { FwbModal } from 'flowbite-vue'
import { ref, reactive, onMounted } from 'vue';
import { toast } from 'vue3-toastify';
import 'vue3-toastify/dist/index.css';

defineProps(['isShowModal', 'showModal' ])


const open = ref(false);
const productData = reactive({
    product_name: '',
    quantity: '',
    price: '',
    description: '',
});


const handleSubmit = async () => {
    try {
        const token = localStorage.getItem('token')
        if(!token){
            console.log('Error: No token')
        }

        const response = await fetch('/api/offices/add', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${token}`,
            },
            body: JSON.stringify(productData),
        });

        if (response.status === 201) {
            toast('Successfully registered')
            productData.product_name = ''
            productData.price = ''
            productData.quantity = ''
        } else {
            const res = await response.json()
            console.log(res.error)
            toast('Error in registration')

        }
    } catch (error) {
        alert('An unexpected error occurred')
        console.error(error.message)
    }
};

</script>

<template>

    <fwb-modal
        v-if="isShowModal" @close="showModal" size="lg"
        class="fixed inset-x-0 inset-y-0 z-50 flex items-center justify-center pl-56">
        <template #header>
            <h2>Add Product</h2>
        </template>
        <template #body>
            <div class="w-full space-y-1 ">
                <form @submit.prevent="handleSubmit">
                    <div class="space-y-4">
                        <div>
                            <input
                                v-model="productData.product_name"
                                type="text"
                                placeholder="Product name"
                                class="w-full px-6 py-2 border-2 border-blue-100 rounded-2xl"
                            />
                        </div>

                        <div>
                            <input
                                v-model="productData.quantity"
                                type="text"
                                placeholder="Total Quantity"
                                class="w-full px-6 py-2 border-2 border-blue-100 rounded-2xl"
                            />
                        </div>
                        <div>
                            <input
                                v-model="productData.price"
                                type="text"
                                placeholder="Unit Price"
                                class="w-full px-6 py-2 border-2 border-blue-100 rounded-2xl"
                            />
                        </div>

                        <div>
                            <QuillEditor
                                theme="snow"
                                v-model="productData.description "
                                contentType="html"
                                placeholder="Description ..."
                                class="w-full p-2 bg-white" />
                        </div>
                    </div>

                    <div class="flex justify-between pt-3 space-x-4">
                        <button
                            @click="showModal"
                            type="button"
                            class="px-4 py-1 font-medium text-white bg-red-500 rounded-lg hover:bg-red-600"
                            >
                            Close
                        </button>
                        <button
                            type="submit"
                            class="px-5 py-1 font-medium text-white bg-blue-500 rounded-lg hover:bg-blue-600">
                            Add
                        </button>
                    </div>
                </form>
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

