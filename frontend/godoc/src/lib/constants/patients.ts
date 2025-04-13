import { writable } from 'svelte/store';

// Patient data interface
export interface Patient {
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

// Dummy patient data
export const DUMMY_PATIENTS: Patient[] = [
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

// CSS Classes
export const BUTTON_CLASSES = {
  PRIMARY: "bg-blue-600 hover:bg-blue-700 text-white py-2 px-4 rounded-md",
  PRIMARY_WITH_ICON: "bg-blue-600 hover:bg-blue-700 text-white py-2 px-4 rounded-md flex items-center",
  SECONDARY: "px-4 py-2 border border-gray-300 rounded-md text-gray-700 hover:bg-gray-50",
  CLOSE: "text-gray-500 hover:text-gray-700"
};

export const FORM_CLASSES = {
  CONTAINER: "space-y-4",
  INPUT: "mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500",
  LABEL: "block text-sm font-medium text-gray-700"
};

export const LAYOUT_CLASSES = {
  PAGE_CONTAINER: "max-w-3xl mx-auto p-4",
  SECTION_CONTAINER: "bg-white shadow-md rounded-lg overflow-hidden",
  SECTION_CONTENT: "p-6"
};

// Create a store for patients data
export const patientsStore = writable<Patient[]>([...DUMMY_PATIENTS]);

// Helper functions to interact with the store
export const patientStoreActions = {
  addPatient: (patient: Patient) => {
    patientsStore.update(patients => [...patients, patient]);
  },

  updatePatient: (updatedPatient: Patient) => {
    patientsStore.update(patients =>
      patients.map(patient =>
        patient.id === updatedPatient.id ? updatedPatient : patient
      )
    );
  },

  getPatientById: (id: string): Patient | undefined => {
    let result: Patient | undefined;
    patientsStore.subscribe(patients => {
      result = patients.find(patient => patient.id === id);
    })();
    return result;
  }
};

