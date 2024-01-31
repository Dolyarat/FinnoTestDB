<script>
export default {
    props:["Person"],
    data(){
        return{
            updatePerson: {
                FirstName: this.Person.FirstName,
                LastName: this.Person.LastName,
                Tel: this.Person.Tel
            },
            isVisible:true
        }
    },
    methods:{
        changeVisible(){
            this.isVisible = !this.isVisible
        },
        
        async putPerson(){
            try {
                const res = await fetch(`http://localhost:8888/persons/${this.Person.PersonID}`, {
                    method: 'PUT',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify(this.updatePerson)
                })
                this.changeVisible()
                location.reload()
            } catch (error){
                console.log(error)
                this.changeVisible()
            }
        },
        async deletePerson(){
            try {
                await fetch(`http://localhost:8888/persons/${this.Person.PersonID}`, {method: 'DELETE'})
                location.reload()
            } catch (error){
                console.log(error)
            }
        }
    },
}

</script>

<template>
    <div class="card">
        <div v-show="isVisible">
            <p>Name: {{ Person.FirstName }} {{ Person.LastName }}</p>
            <p>Tel: {{ Person.Tel }}</p>
            <button @click="changeVisible">edit</button>
            <button @click="deletePerson">delete</button> 

        </div>

        <div v-show="!isVisible">
            <form @submit.prevent="putPerson">
                <label>FirstName: </label>
                <input type="text" v-model="updatePerson.FirstName" /><br>
                <label>LastName: </label>
                <input type="text" v-model="updatePerson.LastName" /><br>
                <label>Tel: </label>
                <input type="text" v-model="updatePerson.Tel" /><br>
                <button @click.prevent="changeVisible">Back</button>
                <button type="submit">Submit</button>

            </form>
            
        </div>
           
    </div>
    
</template>

<style scoped>
.card{
    width: auto;
    height: 120px;
    background-color: peachpuff;
    padding: 10px;
    border-block-end: solid;
    border-block-start: solid;
    text-align: center;

}
input[type=text] {
    width:300px;
    margin-block-end: 10px;
}

</style>