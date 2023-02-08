<script>
    // @ts-nocheck

    // =============
    // # Variables
    export let data = [
        { description: "Learn Go language", priority: 1 },
        { description: "Participate in Open Source with Go", priority: 3 },
        { description: "Create sample project with Go", priority: 2 },
    ];
    let newTodo = {
        description: "",
        priority: 5,
    };
    export let todoAdded = false;

    // =============
    // # Functions
    function insertTodo(newTodo) {
        data.push(newTodo);
        data = data;
        console.log("Todo added: " + newTodo);
        todoAdded = true;
        // newTodo.description = "";
        // newTodo.priority = 0;
        clearNewTodo();
    }

    function clearNewTodo() {
        newTodo = {
            description: "",
            priority: 5,
        };
    }
</script>

<div>
    <h3>New Todo</h3>

    <label>
        <span>Description: </span>
        <textarea bind:value={newTodo.description} />
        <br />
        <span>Priority: {newTodo.priority} </span>
        <input type="range" bind:value={newTodo.priority} min="1" max="5" />
        <br />
        <button on:click={insertTodo(newTodo)}> + </button>
    </label>

    {#if todoAdded}
        <div class="success-message">Todo added!</div>
    {/if}
</div>

<div>
    {#each data as { description, priority }, i}
        <li>
            {i + 1}: {description} : {priority}
        </li>
    {/each}
</div>

<style>
    .success-message {
        color: greenyellow;
        padding: 10px;
    }

    textarea {
        width: 50%;
        height: 30px;
    }
</style>
