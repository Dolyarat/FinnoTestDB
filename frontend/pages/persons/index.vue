<script>
import personCard from '../../components/personCard.vue'
export default{
  components: { personCard },
    async setup(){
        const {data: data} = await useFetch(`http://localhost:8888/persons`)
        return{
            data,
        }
    },
    
}
</script>

<template>
    <div class="container">
        <div class="postForm">
            <h1><u>Add Person</u></h1>
            <postForm/>
        </div>

        <div class="personCard">
            <h1><u>List Person</u></h1>
            <personCard
                v-for="(item, index) in data"
                :key="index"
                :Person="item"
            
            />
        </div>
        
    </div>
</template>

<style scoped>
.container {
    display: grid;
    grid-template-columns: 50% 50%;
    grid-template-areas: "post card";
    
}
.postForm{
    grid-area: post;
    border-right: 5px solid;
    text-align: center;
}
.personCard{
    grid-area: card;
    text-align: center;
}


</style>