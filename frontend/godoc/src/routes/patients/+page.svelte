<script lang="ts">
	// Patient data interface
	interface Patient {
		name: string;
		time: string;
		mobileNumber: string;
	}

	// Dummy patient data
	const patients: Patient[] = [
		{ name: 'John Smith', time: '09:30 AM', mobileNumber: '(555) 123-4567' },
		{ name: 'Sarah Johnson', time: '10:15 AM', mobileNumber: '(555) 234-5678' },
		{ name: 'Robert Williams', time: '11:45 AM', mobileNumber: '(555) 345-6789' },
		{ name: 'Emily Brown', time: '01:30 PM', mobileNumber: '(555) 456-7890' },
		{ name: 'Michael Davis', time: '02:45 PM', mobileNumber: '(555) 567-8901' }
	];
</script>

<div class="max-w-3xl mx-auto p-4">
	<h1 class="text-2xl font-bold mb-6">Patient Records</h1>

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
						<td class="py-3 px-4 border-b border-gray-200">{patient.name}</td>
						<td class="py-3 px-4 border-b border-gray-200">{patient.time}</td>
						<td class="py-3 px-4 border-b border-gray-200">{patient.mobileNumber}</td>
					</tr>
				{/each}
			</tbody>
		</table>
	</div>
</div>

