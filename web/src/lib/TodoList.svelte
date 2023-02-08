<script>
    // @ts-nocheck
    import { createEventDispatcher } from 'svelte';
    
    const dispatch = createEventDispatcher();
    
    function sayHello() {
		dispatch('welcoming', {
			message: 'Hello from inside TodoList svelte page!'
		});
	}

    export let data = [
        "Learn Go language",
        "Participate in Open Source with Go",
        "Create sample project with Go",
    ];
    let columns = ["Todo"];
    let newRow = [...columns];

    function addRow() {
        data = [...data, [...newRow]];
        newRow = columns;
    }
    function deleteRow(rowToBeDeleted) {
        data = data.filter((row) => row != rowToBeDeleted);
    }

    function insertTodo(newTodo) {
        data.push(newTodo);
        data = data;
        console.log("todos now are: " + data);
    }
</script>

<h3>Todos List</h3>

<button on:click={sayHello}>
	Click to say hello
</button>

<table>
    <tr>
        {#each columns as column}
            <th>{column}</th>
        {/each}
        <th>Actions</th>
    </tr>

    {#each data as row}
        <tr>
            <td contenteditable="true" bind:innerHTML={row} />
            <button on:click={() => deleteRow(row)}>X</button>
        </tr>
    {/each}
</table>
