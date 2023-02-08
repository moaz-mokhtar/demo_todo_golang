<script>
    // @ts-nocheck

    // =============
    // # Variables
    export let data = [
        "Learn Go language",
        "Participate in Open Source with Go",
        "Create sample project with Go",
    ];
    let editedTodo = "Participate in open source Go projects 2";
    export let todoAdded = false;

    // =============
    // # Functions
    function insertTodo(newTodo) {
        data.push(newTodo);
        data = data;
        console.log("Todo added: " + newTodo);
        todoAdded = true;
        newTodo = "";
    }
</script>

<div>
    <h3>Edit Todo</h3>

    <label>
        <span>Description: </span>
        <input type="text" bind:value={editedTodo} />
        <button on:click={insertTodo(editedTodo)}> Apply </button>
    </label>

    {#if todoAdded}
        <div class="success-message">Todo added!</div>
    {/if}
</div>

<style>
    .success-message {
        color: greenyellow;
        padding: 10px;
    }
</style>
