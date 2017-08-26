package main

import (
	"github.com/nullrocks/go-bst/bst"
	"fmt"
)

func main() {
	tree := bst.NewTree()

	nums := []int{99603, 13240, 42451, 41600, 79407, 19518, 42570, 42046, 29805, 37094, 91392, 80159, 31313, 43556, 26087, 16116, 87813, 61766, 94777, 74568, 65400, 92893, 1542, 72709, 21754, 37363, 81316, 94372, 56451, 76894, 24943, 31087, 3107, 70073, 68076, 86284, 14597, 31894, 93135, 53734, 89563, 29679, 26271, 44243, 9347, 94492, 61615, 72509, 44962, 39498, 79143, 61579, 7969, 32114, 93720, 14110, 52119, 80813, 41095, 45794, 83283, 19858, 32454, 47255, 28036, 1730, 73837, 60615, 83152, 73472, 24303, 60309, 97177, 85620, 31437, 61219, 11525, 75431, 86349, 466, 32219, 1979, 17087, 84980, 43088, 66559, 94713, 26944, 46542, 35017, 98812, 56020, 14368, 16814, 11738, 10028, 10533, 91633, 83443, 38535, 94076, 33279, 60949, 42834, 18469, 119, 60055, 36542, 2654, 75123, 17892, 18277, 33, 96934, 43128, 16073, 6917, 42589, 61531, 14698, 76716, 58364, 93423, 90922, 72930, 40019, 67708, 4149, 70542, 89865, 81351, 67528, 83368, 45624, 94948, 23578, 29053, 90812, 24713, 86920, 91083, 49986, 37952, 53563, 72920, 66861, 88397, 8423, 65086, 65549, 91823, 71703, 29585, 95752, 78800, 93099, 84745, 43966, 37519, 50451, 70093, 9245, 28851, 41087, 91169, 24699, 18855, 22371, 84893, 9260, 89031, 57559, 11778, 71003, 84533, 68774, 2673, 13129, 79292, 80119, 23548, 37437, 40960, 51239, 44719, 40644, 39861, 62587, 80144, 27441, 64133, 70555, 37925, 20466, 44108, 11827, 55993, 45780, 72604, 51128, 6044, 36355, 58524, 78446, 13513, 97726, 21648, 17215, 44456, 12311, 1272, 93767, 62832, 91545, 88167, 71255, 55404, 81068, 70219, 6181, 42312, 80322, 37797, 39239, 62794, 2433, 35401, 39109, 36078, 24552, 39122, 23752, 26962, 99671, 93729, 37765, 30154, 25938, 65980, 68177, 32304, 74694, 90041, 34326, 68545, 66305, 30784, 50962, 63530, 38248, 72928, 71558, 34904, 3671, 59739, 71321, 14885, 9597, 92705, 64031, 39925, 23927, 64253, 7762, 6397, 95743, 87074, 2829, 29614, 93511, 56043, 36956, 14681, 48429, 96020, 78427, 18278, 4471, 79881, 54894, 74715, 29704, 35495, 72782, 62904, 30712, 94572, 24371, 70202, 43934, 31933, 54306, 76861, 86175, 86123, 59991, 78965, 7490, 18417, 82462, 49227, 95906, 28543, 40263, 23498, 7423, 30788, 27178, 51582, 14137, 84515, 6504, 59140, 25936, 14507, 402, 92128, 45013, 31347, 67159, 74928, 48220, 84511, 26729, 5155, 3084, 42353, 62025, 70069, 62724, 15685, 11134, 29, 18167, 18507, 9935, 24194, 64581, 69150, 39939, 40311, 95258, 85613, 27025, 36426, 29247, 58643, 9333, 83250, 93551, 96937, 81575, 79657, 31813, 41158, 18105, 97253, 51869, 1064, 22215, 3272, 90297, 40954, 8242, 38611, 83496, 66977, 82240, 6050, 93871, 53672, 45389, 69935, 47619, 85319, 79501, 29968, 48115, 39947, 62316, 76300, 45877, 48180, 87805, 46634, 7312, 32619, 18796, 66357, 34016, 23732, 26666, 3545, 31173, 77583, 48594, 21591, 35158, 16034, 61324, 29179, 13413, 11927, 5312, 42000, 81633, 78554, 27170, 65047, 31697, 24053, 37320, 97368, 16566, 89511, 20927, 45484, 34134, 30366, 99959, 31213, 46973, 37585, 24188, 20874, 66505, 76806, 19037, 83231, 71382, 26655, 19957, 9462, 94113, 82005, 39957, 32507, 65924, 50471, 75236, 81042, 97485, 43920, 5457, 24557, 54664, 16791, 8558, 55018, 18869, 5852, 94536, 29149, 32500, 66429, 6282, 13078, 61451, 22032, 51762, 35996, 65041, 64817, 36339, 23489, 7095, 593, 65105, 81046, 98317, 26716, 62074, 84231, 92406, 87500, 90617, 10602, 41785, 99448, 87281, 17020, 87777, 86077, 17290, 4769, 66284, 15647, 83838, 1611, 36398, 12986, 72190, 46186, 93249, 42521, 87614, 23612, 64797, 24779, 23279, 38146, 39921, 55090, 24877, 14883, 80649, 579, 10998, 32747, 7829, 51648, 87763, 31889, 24420, 11107, 19711, 52989, 97724, 45036, 68716, 55300, 73928, 16119, 63472, 92126, 34815, 54952, 61662, 45715, 48810, 42423, 36360, 45145, 75004, 79045, 14315, 12226, 9428, 15788, 49084, 76628, 48044, 81253, 4433, 98876, 23264, 90085, 13558, 57041, 10468, 56439, 33393, 69591, 23825, 27400, 79139, 98355, 3694, 61306, 68503, 77545, 67426, 52140, 35616, 52316, 37352, 94775, 5112, 61683, 15760, 13238, 11232, 59080, 48368, 73614, 16879, 182, 55527, 98802, 44387, 55506, 52021, 54468, 20034, 99927, 12843, 84399, 91262, 57143, 368, 2127, 70004, 63430, 44074, 94658, 38225, 77713, 61164, 26685, 15309, 74089, 78867, 90144, 95459, 95648, 69439, 18714, 73595, 3434, 64413, 91784, 72460, 93202, 52828, 20161, 29397, 31158, 83114, 46251, 12279, 71781, 29417, 89635, 22928, 33098, 92455, 2320, 59056, 90822, 76863, 78700, 92897, 22225, 1578, 67336, 17022, 69473, 47261, 43668, 45344, 80788, 28720, 99700, 51524, 45545, 28545, 15016, 80107, 60055, 68930, 36625, 40004, 23890, 70719, 52948, 32484, 27138, 7731, 40258, 45976, 50988, 95511, 45301, 40838, 52302, 26332, 16265, 92491, 9397, 14680, 97290, 86474, 92058, 30489, 46208, 71871, 55315, 74342, 85707, 87398, 14006, 97942, 38770, 98624, 23885, 97247, 97788, 3456, 10426, 67018, 18151, 55899, 30348, 11662, 8213, 27273, 8964, 43902, 10427, 59261, 99386, 52354, 1299, 47644, 56743, 18694, 42144, 8004, 96790, 93763, 73945, 56986, 5489, 4892, 78578, 41269, 51150, 66133, 87960, 20126, 94582, 81061, 57328, 83897, 11188, 99375, 5693, 15559, 50927, 11724, 4563, 83345, 65836, 42823, 69995, 96798, 15201, 14618, 43041, 74369, 21324, 72826, 65778, 71241, 33376, 45427, 3615, 25876, 33054, 76862, 2347, 83152, 55736, 64310, 71332, 36915, 77552, 51443, 54684, 59800, 39824, 90134, 55833, 84813, 79945, 18682, 63474, 60203, 5157, 60822, 66450, 68238, 59114, 46937, 50768, 21604, 72678, 41010, 87740, 42155, 43332, 45369, 59423, 86147, 12199, 83456, 17899, 1694, 81632, 94517, 50927, 67966, 35905, 48121, 46816, 67339, 34951, 96570, 45916, 69450, 80890, 84611, 75741, 70070, 93809, 4454, 35460, 24256, 91334, 60569, 41857, 42597, 45095, 76486, 47662, 82622, 32183, 71989, 31001, 38857, 19090, 22863, 91610, 98776, 29725, 83733, 54101, 19836, 69962, 10587, 30034, 9773, 65735, 98348, 59748, 73977, 26810, 97066, 11397, 98541, 15162, 34220, 8358, 21731, 20806, 6654, 69747, 955, 98370, 52021, 15713, 81764, 61552, 56270, 84781, 51248, 24718, 14771, 76975, 3656, 63366, 37896, 85245, 97254, 62439, 78070, 84297, 68954, 39511, 65119, 81534, 95073, 28589, 4736, 21113, 77074, 63441, 78571, 31475, 34588, 93515, 36488, 13392, 99229, 42535, 78295, 96584, 52707, 2398, 14649, 41600, 33885, 26729, 58624, 3882, 17709, 11374, 59542, 37044, 99890, 21384, 11707, 24433, 25850, 83335, 27814, 74394, 57805, 12080, 27910, 16672, 59905, 88491, 67970, 92617, 77137, 61167, 2549, 23321, 77662, 82790, 85828, 89022, 5783, 12759, 52278, 83643, 97217, 54723, 1737, 75599, 71217, 41595, 90567, 76116, 45397, 56580, 20191, 36844, 80808, 59337, 10615, 53049, 87995, 92305, 11843, 32829, 28576, 83167, 88732, 95749, 45204, 66510, 77646, 68764, 61370, 46009, 78795, 21516, 42320, 97505, 50970, 20549, 4338, 32770, 61415, 55093, 18350, 45163, 15321, 38560, 3854, 1847, 78671, 45167, 53449, 46140, 94751, 73760, 46463, 35108, 73618, 42361, 40816, 82077, 49140, 58036, 58545, 2769, 92491, 32700, 70175, 78590, 66614, 36759, 21297, 81830, 32005, 55543, 69689}

	fmt.Println(len(nums))

	for _, n := range (nums) {
		tree.Insert(n)
	}

	//tree.Remove(10)
	//
	sum := 0

	for _, n := range (nums) {
		if tree.Remove(n) {
			sum += 1
		}
	}

	fmt.Println(sum)

	//tree.Insert(bst.NewNode(11))
	//tree.Insert(bst.NewNode(112))
	//tree.Insert(bst.NewNode(13))
	//tree.Insert(bst.NewNode(1))
	//tree.Insert(bst.NewNode(111))
	//
	//fmt.Println(tree.RightFirst())
	//fmt.Println(tree.LeftFirst())
	//fmt.Println(tree.RootFirst())
	//
	//fmt.Println(tree.Remove(13))
	//fmt.Println(tree.Remove(10))
	//fmt.Println(tree.Remove(11))
	//fmt.Println(tree.Remove(111))
	//fmt.Println(tree.Remove(112))
	//fmt.Println(tree.Remove(111))
	//fmt.Println(tree.Remove(1))
	//
	fmt.Println(tree.RightFirst())
	fmt.Println(tree.LeftFirst())
	fmt.Println(tree.RootFirst())

	//x, y, z := 4, 5, 6
	//
	////aX, aY, aZ := *x, *y, *z
	//var aX *int = &x
	//var aY *int = &y
	//var aZ *int = &z
	//
	//fmt.Println(*aX)
	//fmt.Println(aY)
	//fmt.Println(aZ)
	//
	//aX = aZ
	//
	//fmt.Println(*aX)
	//fmt.Println(aY)
	//fmt.Println(aZ)

}
