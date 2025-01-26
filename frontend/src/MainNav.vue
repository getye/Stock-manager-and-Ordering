<script setup>

</script>

<template>
    <nav class="fixed top-0 z-20 w-full p-3 px-10 pr-20 space-x-6 font-semibold text-right text-white bg-blue-400 shadow-xl">
        <router-link to="/" class="hover:underline">Home</router-link>
        <router-link to="/user" class="hover:underline">Users</router-link>
        <router-link to="/product" class="hover:underline">Products</router-link>
        <router-link to="/order" class="hover:underline">Orders</router-link>
        <router-link to="/about" class="hover:underline">about</router-link>
        <router-link to="/contact" class="hover:underline">contact</router-link>
        <router-link to="/login" class="hover:underline">login</router-link>
    </nav>
</template>
