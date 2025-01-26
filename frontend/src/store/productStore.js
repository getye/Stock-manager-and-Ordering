import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export const useProductStore = defineStore('product', () => {
    const storedProducts = ref([])
    const searchTerm = ref('');

    function setStoredProducts(office){
        storedProducts.value = office
    }

    function setSearchTerm(term) {
        searchTerm.value = term;
    }

    const filteredProducts = computed(() => {
        const term = searchTerm.value.toLowerCase();
        return storedProducts.value.filter(
        (product) =>
            product.product_name?.toLowerCase().includes(term) 
        )
    })

    return {
        storedProducts, setStoredProducts,
        searchTerm, setSearchTerm,
        filteredProducts,
    }
  })
