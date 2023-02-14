<script>
    import { onMount } from "svelte";

    const PORT_SERVER = 3000;
    const SERVER = "http://localhost:" + PORT_SERVER;
    console.log("SERVER: " + SERVER);

    let list = [];
    let todoItem = clearItem();

    onMount(async () => {
        await getAll();
    });

    function clearItem() {
        return {
            id: -1,
            description: "",
            priority: 5,
        };
    }

    async function getAll() {
        console.log("getAll called");
        const ENDPOINT = SERVER + "/todos";
        console.log("ENDPOINT: " + ENDPOINT);

        const response = await fetch(ENDPOINT, {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
            },
        })
            .then((response) => response.json())
            .then((jsonResp) => {
                console.log("jsonResp: " + jsonResp);
                list = jsonResp.map((item) => {
                    console.log("Type of item: " + typeof item);
                    console.log(item.id);
                    console.log(item.description);
                    return item;
                });
                console.log("list after stringfy: " + list);
            });
        console.log(list);
    }

    async function addTodo(todo) {
        console.log("addTodo called: " + todo);

        if (todo.description.trim() == "") {
            const message = "Todo description is empty !!!";
            alert(message);
            return Error(message);
        }
        if (todo.priority < 0 && todo.priority > 5) {
            const message = "Kindly confirm that todo priority between 1-5 !!!";
            alert(message);
            return Error(message);
        }

        const ENDPOINT = SERVER + "/todo";
        console.log("ENDPOINT: " + ENDPOINT);

        const response = await fetch(ENDPOINT, {
            method: "POST",
            body: JSON.stringify(todo),
            headers: {
                "Content-Type": "application/json",
            },
        });
        console.log("Response: " + response);
        let newId = await response.json();
        console.log("New todo id: " + newId);
        await getAll();
        todoItem = clearItem();
    }

    async function deleteTodo(todoId) {
        console.log("deleteTodo called: " + todoId);
        const ENDPOINT = SERVER + "/todo/" + todoId;
        console.log("ENDPOINT: " + ENDPOINT);

        const response = await fetch(ENDPOINT, {
            method: "DELETE",
            headers: {
                "Content-Type": "application/json",
            },
        });
        console.log("Response: " + response);
        let feedback = await response.json();
        console.log("Feedback: " + feedback);

        await getAll();
        // todoItem = clearItem();
    }

    async function updateTodo(todo) {
        console.log("updateTodo called: " + todo);

        if (todo.description.trim() == "") {
            const message = "Todo description is empty !!!";
            alert(message);
            return Error(message);
        }
        if (todo.priority < 0 && todo.priority > 5) {
            const message = "Kindly confirm that todo priority between 1-5 !!!";
            alert(message);
            return Error(message);
        }

        const ENDPOINT = SERVER + "/todo/" + todo.id;
        console.log("ENDPOINT: " + ENDPOINT);

        const response = await fetch(ENDPOINT, {
            method: "PUT",
            body: JSON.stringify(todo),
            headers: {
                "Content-Type": "application/json",
            },
        });
        console.log("Response: " + response);
        let feedback = await response.json();
        console.log("Feedback: " + feedback);

        await getAll();
        // todoItem = clearItem();
    }
</script>

<div class="container">
    <h3>New Todo</h3>

    <div style="display: flex; flex-direction: row; align-items: center;">
        <div style="flex: 1;">
            <input
                style="width:90%"
                type="text"
                placeholder="Describe your Todo ..."
                id="description"
                bind:value={todoItem.description}
            />
        </div>
        <div style="flex: 1;">
            <label for="priortiy">Priority: {todoItem.priority} </label>
            <input
                id="priority"
                type="range"
                bind:value={todoItem.priority}
                min="1"
                max="5"
            />
        </div>
        <div style="flex; 1;">
            <button on:click={() => addTodo(todoItem)}> + </button>
        </div>
    </div>
</div>

<div class="container">
    <h3>Todos list</h3>
    <table>
        <thead>
            <tr>
                <th>Description</th>
                <th>Priority</th>
                <th>Actions</th>
            </tr>
        </thead>
        <tbody>
            {#each list as todo, id}
                <tr>
                    <td
                        contenteditable="true"
                        bind:innerHTML={todo.description}
                    />
                    <td contenteditable="true" bind:innerHTML={todo.priority} />
                    <td>
                        <button on:click={() => updateTodo(todo)}>Edit</button>
                        <button on:click={() => deleteTodo(todo.id)}
                            >Delete</button
                        >
                    </td>
                </tr>
            {/each}
        </tbody>
    </table>
</div>
