<script>
    import Fa from 'svelte-fa';
    import {faPlay, faEraser, faCopy} from '@fortawesome/free-solid-svg-icons';
    import {onMount} from 'svelte';
    import {Modal as BSModal} from 'bootstrap/dist/js/bootstrap.esm';
    import Modal from '../components/Modal.svelte';
    import {toastSuccess} from '../services/Toast';
    import {Beautify} from '../../wailsjs/go/main/GoStructTab';

    const clearModalID = 'gostruct-clear-modal';

    let /** @type {BSModal} */ clearModal;

    let gostruct = '';

    onMount(async () => {
        clearModal = new BSModal(document.getElementById(clearModalID));
    });

    function onBeautify() {
        Beautify(gostruct).then((result) => (gostruct = result));
    }

    function onCopy() {
        navigator.clipboard.writeText(gostruct);
        toastSuccess('Copied to clipboard!', true);
    }

    function onClear() {
        gostruct = '';
        clearModal.hide();
    }
</script>

<button type="button" class="btn btn-outline-secondary mt-3 gostruct-copy" on:click={onCopy}>
    <Fa icon={faCopy} />
</button>
<textarea
    class="form-control gostruct-textarea mono"
    placeholder="Place your Go struct string here..."
    bind:value={gostruct}
></textarea>
<div class="row mt-3">
    <div class="col-8">
        <div class="d-grid">
            <button type="button" class="btn btn-primary" on:click={onBeautify}>
                <Fa icon={faPlay} />
                Beautify
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

<Modal id={clearModalID} size="md" title="Clearing the field">
    <div slot="body">Are you sure you want to clear the field?</div>
    <div slot="footer">
        <button type="button" on:click={onClear} class="btn btn-success">
            <Fa icon={faEraser} />
            Yes
        </button>
        &nbsp;
        <button type="button" data-bs-dismiss="modal" class="btn btn-light">No</button>
    </div>
</Modal>

<style>
    .gostruct-copy {
        position: absolute;
        right: 0;
        margin-right: 34px;
    }

    .gostruct-textarea {
        height: calc(100vh - 142px);
        min-height: calc(100vh - 142px);
        max-height: calc(100vh - 142px);
    }
</style>
