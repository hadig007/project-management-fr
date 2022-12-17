<template>
    <div class="box">
       <AddBox/>
    </div>
</template>

<script>
import AddBox from './AddBox.vue';

export default {
    components: { AddBox }
}
</script>

<style scoped>
.box {

    min-width: 300px;
    height: 80%;
    background-color: rgba(255, 255, 255, 0.455);
    border-radius: 8px;
    box-shadow: 0 0 3px 3px rgba(255, 255, 255, 0.05);
    margin: 1rem;
}
</style>