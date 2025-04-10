<script>
    import Fa from 'svelte-fa';
    import {faPlay, faEraser} from '@fortawesome/free-solid-svg-icons';
    import {onMount} from 'svelte';
    import {Modal as BSModal} from 'bootstrap/dist/js/bootstrap.esm';
    import Modal from '../components/Modal.svelte';
    import {Calculate} from '../../wailsjs/go/main/PercentTab';

    const clearModalID = 'percent-clear-modal';

    let /** @type {BSModal} */ clearModal;

    let input1 = '';
    let input2 = '';
    let output = '';

    onMount(async () => {
        clearModal = new BSModal(document.getElementById(clearModalID));
    });

    function onCalculate() {
        Calculate(input1, input2).then((/** @type {string} */ result) => (output = result));
    }

    function onClear() {
        input1 = '';
        input2 = '';
        output = '';
        clearModal.hide();
    }
</script>

<!-- svelte-ignore a11y-label-has-associated-control -->
<label class="form-label"><b>Percentage difference</b></label>
<div class="mb-3">
    <input type="text" class="form-control" bind:value={input1} placeholder="Enter first value" />
</div>
<div class="mb-3">
    <input type="text" class="form-control" bind:value={input2} placeholder="Enter second value" />
</div>
<div class="row mb-3">
    <div class="col-8">
        <div class="d-grid">
            <button type="button" class="btn btn-primary" on:click={onCalculate}>
                <Fa icon={faPlay} />
                Calculate
            </button>
        </div>
    </div>
    <div class="col-4">
        <div class="d-grid">
            <button type="button" class="btn btn-light" on:click={() => clearModal.show()}>
                <Fa icon={faEraser} />
                Clear
            </button>
        </div>
    </div>
</div>
<div class="text-success">
    <b>{output}</b>
</div>

<Modal id={clearModalID} size="md" title="Clearing the fields">
    <div slot="body">Are you sure you want to clear the fields?</div>
    <div slot="footer">
        <button type="button" on:click={onClear} class="btn btn-success">
            <Fa icon={faEraser} />
            Yes
        </button>
        &nbsp;
        <button type="button" data-bs-dismiss="modal" class="btn btn-light">No</button>
    </div>
</Modal>
