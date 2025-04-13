<script lang="ts">
	import {
		type Patient,
		BUTTON_CLASSES,
		FORM_CLASSES,
		LAYOUT_CLASSES,
		patientsStore,
		patientStoreActions
	} from '$lib/constants/patients';

	// Subscribe to the store
	let patients: Patient[] = $state([]);
	patientsStore.subscribe((value) => {
		patients = value;
	});

	// Form state
	let newPatient = $state({
		name: '',
		time: '',
		date: '',
		timeInput: '',
		mobileNumber: '',
		address: '',
		remainingFees: 0,
		diagnostics: '',
		allergies: '',
		existingIssues: ''
	});

	let isModalOpen = $state(false);
</script>

<div class={LAYOUT_CLASSES.PAGE_CONTAINER}>
	<div class="flex justify-between items-center mb-6">
		<h1 class="text-2xl font-bold">Patient Records</h1>
		<button onclick={() => (isModalOpen = true)} class={BUTTON_CLASSES.PRIMARY_WITH_ICON}>
			<svg
				xmlns="http://www.w3.org/2000/svg"
				class="h-5 w-5 mr-2"
				viewBox="0 0 20 20"
				fill="currentColor"
			>
				<path
					fill-rule="evenodd"
					d="M10 5a1 1 0 011 1v3h3a1 1 0 110 2h-3v3a1 1 0 11-2 0v-3H6a1 1 0 110-2h3V6a1 1 0 011-1z"
					clip-rule="evenodd"
				/>
			</svg>
			Add Patient
		</button>
	</div>

	<div class="overflow-x-auto">
		<table class="w-full border-collapse">
			<thead>
				<tr>
					<th class="py-3 px-4 text-left bg-gray-100 font-semibold border-b border-gray-200"
						>Name</th
					>
					<th class="py-3 px-4 text-left bg-gray-100 font-semibold border-b border-gray-200"
						>Appointment Time</th
					>
					<th class="py-3 px-4 text-left bg-gray-100 font-semibold border-b border-gray-200"
						>Mobile Number</th
					>
				</tr>
			</thead>
			<tbody>
				{#each patients as patient}
					<tr class="hover:bg-gray-50">
						<td class="py-3 px-4 border-b border-gray-200"
							><a href={`/patients/${patient.id}`} class="text-blue-600 hover:underline"
								>{patient.name}</a
							></td
						>
						<td class="py-3 px-4 border-b border-gray-200">{patient.time}</td>
						<td class="py-3 px-4 border-b border-gray-200">{patient.mobileNumber}</td>
					</tr>
				{/each}
			</tbody>
		</table>
	</div>
</div>

{#if isModalOpen}
	<div class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
		<div class="bg-white p-6 rounded-lg shadow-xl w-full max-w-md">
			<div class="flex justify-between items-center mb-4">
				<h2 class="text-xl font-bold">Add New Patient</h2>
				<button
					aria-label="close"
					onclick={() => (isModalOpen = false)}
					class={BUTTON_CLASSES.CLOSE}
				>
					<svg
						xmlns="http://www.w3.org/2000/svg"
						class="h-6 w-6"
						fill="none"
						viewBox="0 0 24 24"
						stroke="currentColor"
					>
						<path
							stroke-linecap="round"
							stroke-linejoin="round"
							stroke-width="2"
							d="M6 18L18 6M6 6l12 12"
						/>
					</svg>
				</button>
			</div>

			<form
				onsubmit={(e) => {
					e.preventDefault();

					// Format time from date and time inputs
					let formattedTime = '';
					if (newPatient.date && newPatient.timeInput) {
						const dateObj = new Date(`${newPatient.date}T${newPatient.timeInput}`);
						formattedTime = dateObj.toLocaleString('en-US', {
							hour: 'numeric',
							minute: 'numeric',
							hour12: true
						});
					}

					// Create new patient object with required fields
					const patient: Patient = {
						id: `p${String(patients.length + 1).padStart(3, '0')}`,
						name: newPatient.name,
						time: formattedTime,
						mobileNumber: newPatient.mobileNumber,
						address: newPatient.address || '',
						remainingFees: newPatient.remainingFees || 0,
						diagnostics: newPatient.diagnostics || '',
						allergies: newPatient.allergies || '',
						existingIssues: newPatient.existingIssues || ''
					};

					patientStoreActions.addPatient(patient);
					isModalOpen = false;
				}}
				class="space-y-4"
			>
				<div>
					<label for="name" class="block text-sm font-medium text-gray-700">Name</label>
					<input
						type="text"
						id="name"
						bind:value={newPatient.name}
						required
						class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
					/>
				</div>

				<div>
					<label for="date" class="block text-sm font-medium text-gray-700">Appointment Date</label>
					<div class="mt-1 relative">
						<input
							type="date"
							id="date"
							bind:value={newPatient.date}
							required
							class="block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 text-gray-800 bg-white"
						/>
					</div>
				</div>

				<div>
					<label for="time" class="block text-sm font-medium text-gray-700">Appointment Time</label>
					<div class="mt-1 relative">
						<input
							type="time"
							id="time"
							bind:value={newPatient.timeInput}
							required
							class="block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 text-gray-800 bg-white"
						/>
					</div>
				</div>

				<div>
					<label for="mobile" class="block text-sm font-medium text-gray-700">Mobile Number</label>
					<input
						type="text"
						id="mobile"
						bind:value={newPatient.mobileNumber}
						required
						placeholder="e.g., (555) 123-4567"
						class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
					/>
				</div>

				<div class="flex justify-end space-x-3 pt-4">
					<button
						type="button"
						onclick={() => (isModalOpen = false)}
						class="px-4 py-2 border border-gray-300 rounded-md text-gray-700 hover:bg-gray-50"
					>
						Cancel
					</button>
					<button
						type="submit"
						class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700"
					>
						Add Patient
					</button>
				</div>
			</form>
		</div>
	</div>
{/if}
