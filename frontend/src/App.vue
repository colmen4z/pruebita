<template>
  <div class="min-h-screen bg-gray-100">
    <div class="container mx-auto p-8">
      <h1 class="text-3xl font-bold text-center mb-8">Gestión de Usuarios</h1>

      <div class="bg-white p-6 rounded-lg shadow-md mb-8">
        <h2 class="text-xl font-semibold mb-4">{{ editingUser.id ? 'Editar Usuario' : 'Crear Usuario' }}</h2>
        <form @submit.prevent="submitForm" class="space-y-4">
          <div>
            <label class="block text-sm font-medium text-gray-700">Nombre</label>
            <input v-model="editingUser.name" type="text" required
                   class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500">
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700">Email</label>
            <input v-model="editingUser.email" type="email" required
                   class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500">
          </div>
          <button type="submit"
                  class="bg-blue-500 text-white px-4 py-2 rounded-md hover:bg-blue-600">
            {{ editingUser.id ? 'Guardar Cambios' : 'Crear Usuario' }}
          </button>
          <button v-if="editingUser.id" @click="cancelEdit" type="button"
                  class="ml-2 bg-gray-400 text-white px-4 py-2 rounded-md hover:bg-gray-500">
            Cancelar
          </button>
        </form>
      </div>

      <div class="bg-white p-6 rounded-lg shadow-md">
        <h2 class="text-xl font-semibold mb-4">Usuarios Registrados</h2>

        <div v-if="loading" class="text-center">Cargando usuarios...</div>

        <div v-else>
          <div v-for="user in users" :key="user.id"
               class="border-b border-gray-200 py-4 last:border-b-0 flex justify-between items-center">
            <div>
              <p class="font-medium">{{ user.name }}</p>
              <p class="text-gray-600">{{ user.email }}</p>
              <p class="text-sm text-gray-500">Creado: {{ formatDate(user.created_at) }}</p>
            </div>

            <div class="flex space-x-2">
              <button @click="editUser(user)"
                      class="bg-blue-500 text-white px-3 py-1 rounded text-sm hover:bg-blue-600">
                Editar
              </button>

              <button @click="confirmDelete(user)"
                      class="bg-red-500 text-white px-3 py-1 rounded text-sm hover:bg-red-600">
                Eliminar
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue'

export default {
  name: 'UserManagement',
  setup() {
    const users = ref([])
    const loading = ref(false)
    const API_URL = 'http://localhost:3000/api'
    
    // Un solo objeto para crear y editar
    const editingUser = ref({
      id: null,
      name: '',
      email: ''
    })

    // Cargar usuarios
    const fetchUsers = async () => {
      loading.value = true
      try {
        const response = await fetch(`${API_URL}/users`)
        const data = await response.json()
        if (response.ok) { // Mejorar el manejo de errores
          users.value = data.data
        } else {
          console.error('API Error:', data.message)
        }
      } catch (error) {
        console.error('Fetch error:', error)
        alert('Error al cargar usuarios')
      } finally {
        loading.value = false
      }
    }

    // Unifica la creación y la edición en una sola función
    const submitForm = async () => {
      if (editingUser.value.id) {
        // Lógica de edición
        await updateUser()
      } else {
        // Lógica de creación
        await createUser()
      }
    }

    const createUser = async () => {
      try {
        const response = await fetch(`${API_URL}/users`, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(editingUser.value)
        })
        
        if (response.ok) {
          editingUser.value = { id: null, name: '', email: '' } // Limpiar formulario
          alert('✅ Usuario creado exitosamente')
          fetchUsers()
        } else {
          const errorData = await response.json();
          alert('Error: ' + errorData.message);
        }
      } catch (error) {
        console.error('Error creating user:', error)
        alert('Error al crear usuario')
      }
    }

    const updateUser = async () => {
      try {
        const response = await fetch(`${API_URL}/users/${editingUser.value.id}`, {
          method: 'PUT',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(editingUser.value)
        })

        if (response.ok) {
          editingUser.value = { id: null, name: '', email: '' }; // Limpiar formulario
          alert('✅ Usuario actualizado exitosamente');
          fetchUsers();
        } else {
          const errorData = await response.json();
          alert('Error: ' + errorData.message);
        }
      } catch (error) {
        console.error('Error updating user:', error)
        alert('Error al actualizar usuario')
      }
    }

    const editUser = (user) => {
      // Clona el objeto para evitar la mutación directa de la lista
      editingUser.value = { ...user };
    }

    const cancelEdit = () => {
      editingUser.value = { id: null, name: '', email: '' };
    }

    const confirmDelete = (user) => {
      if (confirm(`¿Estás seguro de que quieres eliminar a ${user.name}?`)) {
        deleteUser(user.id)
      }
    }

    const deleteUser = async (userId) => {
      try {
        const response = await fetch(`${API_URL}/users/${userId}`, {
          method: 'DELETE'
        })
        
        if (response.ok) {
          alert('✅ Usuario eliminado exitosamente')
          fetchUsers()
        } else {
          const errorData = await response.json()
          alert('Error: ' + errorData.message)
        }
      } catch (error) {
        console.error('Error deleting user:', error)
        alert('Error al eliminar usuario')
      }
    }

    const formatDate = (dateString) => {
      return new Date(dateString).toLocaleDateString('es-ES')
    }

    onMounted(() => {
      fetchUsers()
    })

    return {
      users,
      loading,
      editingUser,
      submitForm,
      editUser,
      cancelEdit,
      confirmDelete,
      deleteUser,
      formatDate
    }
  }
}
</script>