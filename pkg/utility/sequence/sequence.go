package sequence

/// Currently Ceres use the simple snowflake with the public IP of the machine which Ceres running on it 
/// TODO: next version need use the distributed senquence implementation 

/// Senquence
/// interface to generate the unique senquence number
type Senquence interface {
	/// Next will return the next senquence
	/// used in the comer logic not only the comer but the profile
	/// used in the bounty logic
	/// used in the disco logic
	Next() uint64
}
