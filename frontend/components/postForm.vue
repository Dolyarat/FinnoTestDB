<script>
export default{
    data(){
        return{
            Person:{FirstName:"", LastName:"", Tel:""},
        }
    },
    methods:{
        async postPerson(){
            try {
                const res = await fetch(`http://localhost:8888/persons`, {
                    method: 'POST',
                    headers:{
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify(this.Person)
                })
                location.reload()
            } catch (error){
                console.log(error)
            }
        }
    }
}
</script>

<template>
    <div class="post">
        <form @submit.prevent="postPerson">
            <label>FirstName: </label>
            <input type="text" v-model="Person.FirstName" /><br>
            <label>LastName: </label>
            <input type="text" v-model="Person.LastName" /><br>
            <label>Tel: </label>
            <input type="text" v-model="Person.Tel" /><br>
            <button type="submit">Submit</button>
        </form>
    </div>
    
</template>

<style scoped>
.post{
    width: auto;
    height: auto;
    padding: 20px;
    text-align: center;
    
}
input[type=text] {
    width:300px;
    margin-block-end: 10px;
}

</style>