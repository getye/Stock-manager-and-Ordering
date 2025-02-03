<script setup>
import { defineProps, watch, watchEffect } from 'vue';
import { FwbModal } from 'flowbite-vue'
import { ref, reactive, onMounted } from 'vue';
import { toast } from 'vue3-toastify';
import 'vue3-toastify/dist/index.css';

defineProps(['isShowModal', 'showModal' ])


const open = ref(false);
const productData = reactive({
    name: '',
    quantity: '',
    price: '',
    image: null,
    description: '',
})


const handleSubmit = async () => {
    try {
        
        const product = new FormData()
        product.append("name", productData.name)
        product.append("description", productData.description)
        product.append("quantity", productData.quantity)
        product.append("price", productData.price)
        product.append("image", productData.image)

        const response = await fetch('http://localhost:5000/api/products/add', {
            method: 'POST',
            body: product,
        });

        if (response.status === 201) {
            const res = await response.json()
            alert(res.message)
            
            productData.name = ''
            productData.price = ''
            productData.quantity = ''
            productData.description =''
            productData.image = null
        } else {
            const res = await response.json()
            console.log(res.error)
            alert('Error in registration')

        }
    } catch (error) {
        alert('An unexpected error occurred')
        console.error(error.message)
    }
}

const triggerImageUpload = () => {
    document.querySelector("input[type='file']").click()
}

const uploadImage = (event) => {
    const file = event.target.files[0]
    if (file) {
        productData.image = file
    }
}



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
                <div>
                    <div class="space-y-4">
                        <div>
                            <input
                                v-model="productData.name"
                                type="text"
                                placeholder="Product name"
                                class="w-full px-6 py-2 border-2 border-blue-100 rounded-2xl"
                            />
                        </div>

                        <div>
                            <input
                                v-model="productData.quantity"
                                type="number"
                                placeholder="Total Quantity"
                                class="w-full px-6 py-2 border-2 border-blue-100 rounded-2xl"
                            />
                        </div>
                        <div>
                            <input
                                v-model="productData.price"
                                type="number"
                                placeholder="Unit Price"
                                class="w-full px-6 py-2 border-2 border-blue-100 rounded-2xl"
                            />
                        </div>

                        <div>
                            <textarea
                                v-model="productData.description"
                                placeholder="Description ..."
                                class="w-full px-6 border-2 border-blue-100 rounded-2xl" />
                        </div>

                        <div>
                            <button @click="triggerImageUpload"
                                class="w-full p-2 mb-2 font-bold text-white bg-gradient-to-r from-blue-400 to-gray-500 rounded-2xl">
                                <font-awesome-icon :icon="['fas', 'upload']" class="pr-2" />
                                {{ productData.image ? productData.image.name : "Product Image" }}
                            </button>
                            <input type="file" name="image" accept="image/*" @change="uploadImage" class="hidden" />
                        </div>
                    </div>

                    <div class="flex justify-end pt-3 space-x-4">
                        <button
                            @click="handleSubmit"
                            class="px-5 py-1 font-medium text-white bg-blue-500 rounded-lg hover:bg-blue-600">
                            Add
                        </button>
                    </div>
                </div>
            </div>
        </template>

    </fwb-modal>
</template>


