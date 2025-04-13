<script lang="ts">
	import { onMount } from 'svelte';

	// Patient data interface
	interface Patient {
		id: string;
		name: string;
		time: string;
		mobileNumber: string;
		address: string;
		remainingFees: number;
		diagnostics: string;
		allergies: string;
		existingIssues: string;
	}

	// For actual implementation, you would fetch this from an API
	// Here we use dummy data that matches the patients list
	const dummyPatients: Patient[] = [
		{
			id: 'p001',
			name: 'John Smith',
			time: '09:30 AM',
			mobileNumber: '(555) 123-4567',
			address: '123 Main St, Anytown',
			remainingFees: 150,
			diagnostics: 'Type 2 Diabetes',
			allergies: 'Penicillin',
			existingIssues: 'Hypertension, High Cholesterol'
		},
		{
			id: 'p002',
			name: 'Sarah Johnson',
			time: '10:15 AM',
			mobileNumber: '(555) 234-5678',
			address: '456 Oak Ave, Somewhere',
			remainingFees: 75,
			diagnostics: 'Arthritis',
			allergies: 'Latex',
			existingIssues: 'Asthma'
		},
		{
			id: 'p003',
			name: 'Robert Williams',
			time: '11:45 AM',
			mobileNumber: '(555) 345-6789',
			address: '789 Pine Blvd, Elsewhere',
			remainingFees: 200,
			diagnostics: 'Hypertension',
			allergies: 'None',
			existingIssues: 'Anxiety, Insomnia'
		},
		{
			id: 'p004',
			name: 'Emily Brown',
			time: '01:30 PM',
			mobileNumber: '(555) 456-7890',
			address: '101 Maple Dr, Nowhere',
			remainingFees: 0,
			diagnostics: 'Migraine',
			allergies: 'Sulfa drugs',
			existingIssues: 'Depression'
		},
		{
			id: 'p005',
			name: 'Michael Davis',
			time: '02:45 PM',
			mobileNumber: '(555) 567-8901',
			address: '202 Cedar Ln, Anywhere',
			remainingFees: 50,
			diagnostics: 'GERD',
			allergies: 'Shellfish',
			existingIssues: 'Back pain, Obesity'
		}
	];

	let patient: Patient | undefined = $state(undefined);
	let editMode = $state(false);
	let patientCopy: Patient | undefined = $state(undefined);
	let patientId: string | undefined = undefined;
	let notFound = $state(false);

	// Get patient ID from URL
	const { data } = $props();
	// Get patient ID from URL
	onMount(() => {
		patientId = data.params.id;
		loadPatient();
	});

	function loadPatient() {
		if (!patientId) return;

		// In a real app, you would fetch the patient data from an API
		// Here we simulate by finding the patient in our dummy data
		patient = dummyPatients.find((p) => p.id === patientId);

		if (!patient) {
			notFound = true;
		} else {
			notFound = false;
		}
	}

	function startEdit() {
		// Create a deep copy of patient for editing
		patientCopy = patient ? JSON.parse(JSON.stringify(patient)) : undefined;
		editMode = true;
	}

	function cancelEdit() {
		editMode = false;
		patientCopy = undefined;
	}

	function saveChanges() {
		if (!patientCopy) return;

		// In a real app, you would send an API request to update the patient
		// Here we update our local state
		patient = { ...patientCopy };

		// Also update the patient in the dummy list (to persist changes if we navigate away and back)
		const index = dummyPatients.findIndex((p) => p.id === patient?.id);
		if (index !== -1) {
			dummyPatients[index] = { ...patient };
		}

		editMode = false;
	}
</script>

<div class="max-w-3xl mx-auto p-4">
	<div class="mb-6">
		<a href="/patients" class="inline-flex items-center text-blue-600 hover:text-blue-800">
			<svg
				xmlns="http://www.w3.org/2000/svg"
				class="h-5 w-5 mr-1"
				fill="none"
				viewBox="0 0 24 24"
				stroke="currentColor"
			>
				<path
					stroke-linecap="round"
					stroke-linejoin="round"
					stroke-width="2"
					d="M10 19l-7-7m0 0l7-7m-7 7h18"
				/>
			</svg>
			Back to Patient List
		</a>
	</div>

	{#if notFound}
		<div
			class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded relative"
			role="alert"
		>
			<strong class="font-bold">Error!</strong>
			<span class="block sm:inline"> Patient not found.</span>
		</div>
	{:else if patient}
		<div class="bg-white shadow-md rounded-lg overflow-hidden">
			<div class="p-6">
				<div class="flex justify-between items-center mb-6">
					<h1 class="text-2xl font-bold text-gray-800">{patient.name}</h1>
					{#if !editMode}
						<button
							onclick={startEdit}
							class="bg-blue-600 hover:bg-blue-700 text-white py-2 px-4 rounded-md"
						>
							Edit Patient
						</button>
					{/if}
				</div>

				{#if !editMode}
					<div class="grid grid-cols-1 md:grid-cols-2 gap-6">
						<div>
							<h2 class="text-lg font-semibold text-gray-700 mb-2">Contact Information</h2>
							<div class="space-y-2">
								<p><span class="font-medium">Phone:</span> {patient.mobileNumber}</p>
								<p><span class="font-medium">Address:</span> {patient.address}</p>
							</div>
						</div>

						<div>
							<h2 class="text-lg font-semibold text-gray-700 mb-2">Appointment Details</h2>
							<div class="space-y-2">
								<p><span class="font-medium">Last Appointment:</span> {patient.time}</p>
								<p><span class="font-medium">Remaining Fees:</span> ${patient.remainingFees}</p>
							</div>
						</div>
					</div>

					<div class="mt-6">
						<h2 class="text-lg font-semibold text-gray-700 mb-2">Medical Information</h2>
						<div class="space-y-4">
							<div>
								<h3 class="font-medium">Diagnostics</h3>
								<p class="mt-1">{patient.diagnostics || 'None recorded'}</p>
							</div>
							<div>
								<h3 class="font-medium">Allergies</h3>
								<p class="mt-1">{patient.allergies || 'None recorded'}</p>
							</div>
							<div>
								<h3 class="font-medium">Existing Issues</h3>
								<p class="mt-1 whitespace-pre-line">{patient.existingIssues || 'None recorded'}</p>
							</div>
						</div>
					</div>
				{:else if patientCopy}
					<form
						onsubmit={(e) => {
							e.preventDefault();
							saveChanges();
						}}
						class="space-y-4"
					>
						<div>
							<label for="name" class="block text-sm font-medium text-gray-700">Name</label>
							<input
								type="text"
								id="name"
								bind:value={patientCopy.name}
								required
								class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
							/>
						</div>

						<div>
							<label for="mobile" class="block text-sm font-medium text-gray-700"
								>Mobile Number</label
							>
							<input
								type="text"
								id="mobile"
								bind:value={patientCopy.mobileNumber}
								required
								class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
							/>
						</div>

						<div>
							<label for="address" class="block text-sm font-medium text-gray-700">Address</label>
							<input
								type="text"
								id="address"
								bind:value={patientCopy.address}
								required
								class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
							/>
						</div>

						<div>
							<label for="time" class="block text-sm font-medium text-gray-700"
								>Appointment Time</label
							>
							<input
								type="text"
								id="time"
								bind:value={patientCopy.time}
								class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
							/>
						</div>

						<div>
							<label for="remainingFees" class="block text-sm font-medium text-gray-700"
								>Remaining Fees ($)</label
							>
							<input
								type="number"
								id="remainingFees"
								bind:value={patientCopy.remainingFees}
								min="0"
								step="0.01"
								class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
							/>
						</div>

						<div>
							<label for="diagnostics" class="block text-sm font-medium text-gray-700"
								>Diagnostics</label
							>
							<input
								type="text"
								id="diagnostics"
								bind:value={patientCopy.diagnostics}
								class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
							/>
						</div>

						<div>
							<label for="allergies" class="block text-sm font-medium text-gray-700"
								>Allergies</label
							>
							<input
								type="text"
								id="allergies"
								bind:value={patientCopy.allergies}
								class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
							/>
						</div>

						<div>
							<label for="existingIssues" class="block text-sm font-medium text-gray-700"
								>Existing Issues</label
							>
							<textarea
								id="existingIssues"
								bind:value={patientCopy.existingIssues}
								rows="4"
								class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
							></textarea>
						</div>

						<div class="flex justify-end space-x-3 pt-4">
							<button
								type="button"
								onclick={cancelEdit}
								class="px-4 py-2 border border-gray-300 rounded-md text-gray-700 hover:bg-gray-50"
							>
								Cancel
							</button>
							<button
								type="submit"
								class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700"
							>
								Save Changes
							</button>
						</div>
					</form>
				{/if}
			</div>
		</div>
	{:else}
		<div class="flex justify-center items-center h-40">
			<div class="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-blue-600"></div>
		</div>
	{/if}
</div>
