
<script setup>
import { ref, reactive, onMounted, watchEffect, watch } from 'vue';
import { toast } from 'vue3-toastify';
import 'vue3-toastify/dist/index.css';
import { useProductStore } from '../../store/productStore';
import ProductModal from './ProductModal.vue';
import ProductUpdateModal from './ProductUpdateModal.vue';

const productStore = useProductStore()
const products = ref([])
const selectedProduct = ref({})


const isShowModal = ref(false)
const isUpdateModal = ref(false)

const role = localStorage.getItem('userRole')
function showModal () {
    isShowModal.value = !isShowModal.value
}

function updateModal (product) {
    isUpdateModal.value = true
    selectedProduct.value = product
}

function closeUpdateModal () {
    isUpdateModal.value = false
}

watchEffect( async() => {
    if(role === 'Admin'){
        try {
            const token = localStorage.getItem('token')
            if(!token){
                console.log('Error: No token')
            }
            const response = await fetch('/api/offices/view', {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${token}`,
                },
            });

            if(response.status === 200) {
                const data = await response.json()
                products.value = data.data
                productStore.setStoredOffices(data.data)
            }
            else{
                const res = await response.json()
                console.log(res.error)
                alert(res.message)
            }
        } catch (err) {
        console.log(err)
        }
    }else alert("You have no permission")
  })


const handleDeleteOffice = async (product) => {
    if (!window.confirm(`Are you sure you want to delete office: ${product.product_name}?`)) {
        return;
    }

    const token = localStorage.getItem('token')
    if(!token){
        console.log('Error: No token')
    }

    try {
        const response = await fetch(`/api/offices/delete/${product.product_id}`, {
            method: 'DELETE',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${token}`,
            },
        });

        if (response.ok) {
            toast('Office Deleted Successfully')
        } else {
            toast(`Error: ${response.statusText}`)
        }
    } catch (err) {
        toast('An unexpected error occurred')
        console.error(err)
    }
}

watch(()=>{
    if(productStore.filteredProducts.length > 0){
        products.value = productStore.filteredProducts
    }
})
</script>


<template>
    <div>
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
                                :key="product.product_id"
                                class="hover:bg-gray-50">
                                <td class="px-4 py-2 border border-gray-300">{{ product.product_name }}</td>
                                <td class="px-4 py-2 border border-gray-300">{{ product.total_quantity }}</td>
                                <td class="px-4 py-2 border border-gray-300">{{ product.order_quantity }}</td>
                                <td class="px-4 py-2 border border-gray-300">{{ product.total_quantity - product.order_quantity }}</td>
                                <td class="px-4 py-2 border border-gray-300">{{ product.price}}</td>
                                <td class="px-4 py-2 border border-gray-300">{{ product.price * product.total_quantity }}</td>
                                <td class="flex justify-center px-4 py-2 space-x-4 border border-gray-300">
                                    <button
                                        @click="updateModal(product)"
                                        class="text-blue-500">
                                        <font-awesome-icon :icon="['fas', 'pen-to-square']" />
                                    </button>
                                    <button
                                        @click="handleDeleteOffice(product)"
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

