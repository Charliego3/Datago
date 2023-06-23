<script>
    import "uno.css";
    import "@unocss/reset/tailwind.css";

    let splitterDragging = false;

    function toggleSidebar() {
        console.log('toggle sidebar button clicked.');
        splitterDragging = !splitterDragging;
    }

    let count = 1;

	// the `$:` means 're-run whenever these values change'
	$: doubled = count * 2;
	$: quadrupled = doubled * 2;

	function handleClick() {
		count += 1;
	}
</script>

<div flex w-screen h-screen overflow-hidden font-robot divide-x divide="gray/35"
    class="{splitterDragging ? 'cursor-col-resize' : 'cursor-default'}">
    <div relative w-400px h-full flex-none flex flex-col divide-y divide="gray/35">
        <div flex-none h-38px flex items-center text-xl>
            <div w-77px h-full pl-13px pr-11px flex items-center justify-between>
                <div bg-red w-12px h-12px rounded-full />
                <div bg-yellow w-12px h-12px rounded-full />
                <div bg-green w-12px h-12px rounded-full />
            </div>
            <button ml-80px fixed cursor-default px-7px py-1px rounded flex bg="hover:slate/15" on:click={toggleSidebar}>
                <span text="dark:white" i-icons-sidebar-leading opacity-60 cursor-default />
            </button>
        </div>
        <div flex-auto px-3px />
        <div flex-none w-5px absolute mt--1 right--3px h-full cursor-col-resize bg="indigo/0"></div>
    </div>
    <div w-full h-full flex flex-col divide-y divide="gray/35">
        <div flex-none h-38px bg="#FAF9F9 dark:#373736" />
        <div flex-auto px-3px bg="white dark:#2e2e2d">
            <slot />

            <button bg-gray-200 border border-black px-2 rounded shadow-lg on:click={handleClick}>
                Count: {count}
            </button>

            <p>{count} * 2 = {doubled}</p>
            <p>{doubled} * 2 = {quadrupled}</p>
        </div>
    </div>
</div>
