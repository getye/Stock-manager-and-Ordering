<script setup>
import { defineProps, watchEffect } from 'vue';
import { FwbModal } from 'flowbite-vue'
import { ref } from 'vue';

const props = defineProps({
    isUpdateModal: Boolean,
    updateModal: Function,
    selectedProduct: {
        type: Object,
        required: true,
  },
})

const open = ref(false)
const users = ref([])
const productData = ref({})


const handleClose = () => {
    open.value = false
};



const handleUpdateProduct = async () => {
    try {

        const token = localStorage.getItem('token')

        if(!token){
            console.log('Error: No token')
        }

        const response = await fetch(`/api/offices/update/${productData.value.product_id}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${token}`,
            },
            body: JSON.stringify(productData.value),
        });

        if (response.ok) {
            alert('Office Successfully Updated')
            handleClose()
        }
        else {
            alert(`Error: ${response.statusText}`)
        }
    } catch (err) {
        alert('An unexpected error occurred')
        console.error(err)
    }
  };

</script>

<template>
    <fwb-modal v-if="isUpdateModal" @close="updateModal" size="md"
        class="fixed inset-x-0 inset-y-0 z-50 flex items-center justify-center pl-56">
        <template #header>
            <h2>Update Product</h2>
        </template>
        <template #body>
            <div class="w-full max-w-md px-8 space-y-1 ">
                <form @submit.prevent="handleUpdateProduct">
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
                            @click="updateModal"
                            type="button"
                            class="px-4 py-1 font-medium text-white bg-red-500 rounded-lg hover:bg-red-600"
                            >
                            Close
                        </button>
                        <button
                            type="submit"
                            class="px-5 py-1 font-medium text-white bg-blue-500 rounded-lg hover:bg-blue-600"
                            >
                            Update
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

