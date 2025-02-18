<script setup>
import { defineProps, watchEffect } from 'vue';
import { FwbModal } from 'flowbite-vue'
import { ref, watch } from 'vue';

const props = defineProps({
    isUpdateModal: Boolean,
    updateModal: Function,
    selectedProduct: {
        type: Object,
        required: true,
  },
})

const open = ref(false)
const productData = ref()


const handleClose = () => {
    open.value = false
}

watch(()=>{
   productData.value = props.selectedProduct
})

const handleUpdateProduct = async () => {
    try {
        /*
        const token = localStorage.getItem('token')
        if(!token){
            console.log('Error: No token')
        }
            */


        const id = productData.value.id
        const response = await fetch(`http://localhost:5000/api/products/update/${id}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
                //'Authorization': `Bearer ${token}`,
            },
            body: JSON.stringify(productData.value),
        })

        if (response.ok) {
            alert('Product Successfully Updated')
            productData.name = ''
            productData.price = ''
            productData.quantity = ''
            productData.description =''
            productData.image = null
        }
        else {
            alert(`Error: ${response.statusText}`)
        }
    } catch (err) {
        alert('An unexpected error occurred')
        console.error(err)
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
    <fwb-modal v-if="isUpdateModal" @close="updateModal" size="lg"
        class="fixed inset-x-0 inset-y-0 z-50 flex items-center justify-center pl-56">
        <template #header>
            <h2>Update Product</h2>
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
                                {{ productData.image.name? productData.image.name : productData.image }}
                            </button>
                            <input type="file" name="image" accept="image/*" @change="uploadImage" class="hidden" />
                        </div>
                    </div>

                    <div class="flex justify-end pt-3 space-x-4">
                        <button
                            @click="handleUpdateProduct"
                            class="px-5 py-1 font-medium text-white bg-blue-500 rounded-lg hover:bg-blue-600">
                            Update
                        </button>
                    </div>
                </div>
            </div>
        </template>

    </fwb-modal>
</template>



