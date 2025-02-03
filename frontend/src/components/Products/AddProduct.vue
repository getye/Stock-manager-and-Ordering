
<script setup>
import { ref, reactive, onMounted, watchEffect, watch } from 'vue';
import { toast } from 'vue3-toastify';
import 'vue3-toastify/dist/index.css';
import { useProductStore } from '../../store/productStore';
import ProductModal from './ProductModal.vue';
import ProductUpdateModal from './ProductUpdateModal.vue';

import { FwbModal } from 'flowbite-vue'


const productStore = useProductStore()
const products = ref([])
const selectedProduct = ref({})


const isShowModal = ref(false)
const isUpdateModal = ref(false)
const isDetailModal = ref(false)


const role = localStorage.getItem('userRole')
const showModal = () =>{
    isShowModal.value = !isShowModal.value
}

const updateModal = (product) => {
    isUpdateModal.value = true
    selectedProduct.value = product
}

const closeUpdateModal = () => {
    isUpdateModal.value = false
}

const detailModal = (product) => {
    isDetailModal.value = !isDetailModal.value
    selectedProduct.value = product
}

watch( async() => {
    //if(role === 'Admin'){
        try {
            /*
            const token = localStorage.getItem('token')
            if(!token){
                console.log('Error: No token')
            }
                */
            const response = await fetch('http://localhost:5000/api/products/view', {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                    //'Authorization': `Bearer ${token}`,
                },
            });

            if(response.status === 200) {
                const data = await response.json()
                console.log(data.products)
                products.value = data.products
            }
            else{
                const res = await response.json()
                console.log(res.error)
                alert(res.error)
            }
        } catch (err) {
        console.log(err)
        }
    //}else alert("You have no permission")
  })


const handleDeleteProduct = async (product) => {
    if (!window.confirm(`Are you sure you want to delete: ${product.name}?`)) {
        return;
    }

    /*
    const token = localStorage.getItem('token')
    if(!token){
        console.log('Error: No token')
    }
        */

    try {
        const response = await fetch(`http://localhost:5000/api/products/delete/${product.id}`, {
            method: 'DELETE',
            headers: {
                'Content-Type': 'application/json',
                //'Authorization': `Bearer ${token}`,
            },
        });

        if (response.ok) {
            alert('Product Deleted Successfully')
        } else {
            alert(`Error: ${response.statusText}`)
        }
    } catch (err) {
        alert('An unexpected error occurred')
        console.error(err)
    }
}

const getImageUrl = (imagePath) => {
  return `http://localhost:5000/${imagePath}`;
}

watch(()=>{
    if(productStore.filteredProducts.length > 0){
        products.value = productStore.filteredProducts
    }
})
</script>


<template>
    <div class="flex justify-center ">
        <div class="mt-16">
            <div class="p-2">
                <div class="flex items-center justify-start w-full pb-3 space-x-4">
                    <button
                        @click="showModal"
                        class="inline-flex gap-2 px-2 py-1 font-bold text-white bg-blue-400 rounded-lg"
                        >
                        <font-awesome-icon
                            :icon="['fas', 'circle-plus']"
                            class="pt-1"
                            />
                        Add Product
                    </button>
                    <input
                        type="text"
                        v-model="productStore.searchTerm"
                        placeholder="Search Products"
                        class="py-1 pl-4 border-2 border-blue-500 w-60 rounded-2xl"
                    />
                </div>
                <div v-if="products.length > 0">
                    <table class="w-full border border-collapse border-gray-300 table-auto">
                        <thead class="text-white bg-blue-300">
                            <tr>
                                <th class="px-4 py-1 border border-gray-300">Product Name</th>
                                <th class="px-4 py-1 border border-gray-300">Total Quantity</th>
                                <th class="px-4 py-1 border border-gray-300">Order Quantity</th>
                                <th class="px-4 py-1 border border-gray-300">Remaining</th>
                                <th class="px-4 py-1 border border-gray-300">Unit price</th>
                                <th class="px-4 py-1 border border-gray-300">Total Price</th>
                                <th class="px-4 py-1 border border-gray-300">Actions</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr v-for="product in products"
                                :key="product.id"
                                class="hover:bg-gray-50">
                                <td class="px-4 py-2 border border-gray-300">{{ product.name }}</td>
                                <td class="px-4 py-2 border border-gray-300">{{ product.quantity }}</td>
                                <td class="px-4 py-2 border border-gray-300">{{ product.order }}</td>
                                <td class="px-4 py-2 border border-gray-300">{{ product.quantity - product.order }}</td>
                                <td class="px-4 py-2 border border-gray-300">{{ product.price}}</td>
                                <td class="px-4 py-2 border border-gray-300">{{ product.price * product.quantity }}</td>
                                <td class="flex justify-center px-4 py-2 space-x-4 border border-gray-300">
                                    <button
                                        @click="detailModal(product)"
                                        title="View detail"
                                        class="text-green-500"
                                    >
                                        <font-awesome-icon icon="eye" />
                                    </button>
                                    <button
                                        @click="updateModal(product)"
                                        title="Updaate Product"
                                        class="text-blue-500"
                                    >
                                        <font-awesome-icon :icon="['fas', 'pen-to-square']" />
                                    </button>
                                    <button
                                        @click="handleDeleteProduct(product)"
                                        title="Delete this product"
                                        class="text-red-500"
                                    >
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

        <fwb-modal
            v-if="isDetailModal" @close="detailModal" size="lg"
            class="fixed inset-x-0 inset-y-0 z-50 flex items-center justify-center pl-56">
            <template #header>
                <h2>{{ selectedProduct.name }}</h2>
            </template>
            <template #body>
                <div class="flex justify-between w-full gap-2 px-6 py-2">
                    <div>
                        <p> Total Quantity: {{ selectedProduct.quantity }}</p>
                        <p> Unit Price: {{ selectedProduct.price }}</p>
                        <p> Total Orders: {{ selectedProduct.order }}</p>
                    </div>
                    <div>
                        <img
                            :src="getImageUrl(selectedProduct.image)"
                            alt="image"
                        />
                    </div>
                </div>
                <div class="px-3 py-2 border-t-2 border-blue-300">
                    <p>{{ selectedProduct.description }}</p>
                </div>
            </template>

        </fwb-modal>


        <ProductModal
            :isShowModal="isShowModal"
            :showModal="showModal"/>

        <ProductUpdateModal
            :isUpdateModal="isUpdateModal"
            :updateModal="closeUpdateModal"
            :selectedProduct = "selectedProduct"
            />
        </div>
  </template>

