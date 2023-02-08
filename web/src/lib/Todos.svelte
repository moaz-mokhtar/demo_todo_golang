<script>
// @ts-nocheck


	let data=["Learn Go language", "Participate in Open Source with Go", "Create sample project with Go"];
	let newTodo="Participate in open source Go projects 2";
	let columns = ["Todo"]
	let newRow = [...columns];

	function addRow() {
		data = [...data, [...newRow]]
		newRow = columns
	}
	function deleteRow(rowToBeDeleted) {
		data = data.filter(row => row != rowToBeDeleted)
	}

	function insertTodo(newTodo){
		data.push(newTodo);
		data=data;
		console.log("todos now are: "+data);
	}
</script>
<div>
	<p>
		Add new todo said: "{newTodo}"
	</p>
	<textarea bind:value={newTodo}></textarea>
	
	<button on:click={insertTodo(newTodo)}> add </button>
	<p>
		Todos are: {data.join(" || ")}
	</p>
</div>

<table>
	<tr>
		{#each columns as column}
			<th>{column}</th>
		{/each}
	</tr>
	
	{#each data as row}
		<tr>
			{#each row as cell}
			<td contenteditable="true" bind:innerHTML={cell} />
			{/each}
			<button on:click={() => deleteRow(row)}>X</button>
		</tr>
	{/each}
	<tr style="color: grey">
		{#each newRow as column}
			<td contenteditable="true" bind:innerHTML={column} />
		{/each}
		<button on:click={addRow}>add</button>
	</tr>
	<pre style="background: #eee">{JSON.stringify(data, null, 2)}</pre>
</table>
